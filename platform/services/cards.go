package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/sultaniman/confetti/platform/entities"
	"github.com/sultaniman/confetti/platform/http"
	"github.com/sultaniman/confetti/platform/repo"
	"github.com/sultaniman/confetti/platform/schema"
	"github.com/sultaniman/pwc/crypto"
	"github.com/sultaniman/pwc/gen"
)

type CardService interface {
	Generate(options *schema.CardOptions) (*schema.NewCardResponse, error)
	Get(cardId uuid.UUID) (*schema.CardResponse, error)
	List(userId uuid.UUID) ([]schema.CardResponse, error)
	Create(userId uuid.UUID, newCard *schema.NewCardRequest) (*schema.CardResponse, error)
	Update(cardId uuid.UUID, updateRequest *schema.UpdateCardRequest) error
	Delete(cardId uuid.UUID) error
	Decrypt(cardId uuid.UUID) (*schema.PlainCardResponse, error)
	ClaimExists(cardId uuid.UUID, userId uuid.UUID) bool
}

type cardService struct {
	privateKey *rsa.PrivateKey
	cardsRepo  repo.CardRepo
	usersRepo  repo.UserRepo
}

func NewCardService(usersRepo repo.UserRepo, cardsRepo repo.CardRepo, privateKey *rsa.PrivateKey) CardService {
	return &cardService{
		privateKey: privateKey,
		cardsRepo:  cardsRepo,
		usersRepo:  usersRepo,
	}
}

func (c *cardService) Generate(options *schema.CardOptions) (*schema.NewCardResponse, error) {
	_, card, err := gen.GenerateClassicCard(options.IncludeSymbols, options.DigitsArea, false)
	if err != nil {
		return nil, http.InternalError(err)
	}

	return &schema.NewCardResponse{
		Data: string(card.GetBytes()),
		Key:  card.Passphrase,
	}, nil
}

func (c *cardService) Get(cardId uuid.UUID) (*schema.CardResponse, error) {
	card, err := c.cardsRepo.Get(cardId)
	if err != nil {
		return nil, c.handleError(err)
	}

	return c.cardToResponse(card), nil
}

func (c *cardService) List(userId uuid.UUID) ([]schema.CardResponse, error) {
	cards, err := c.cardsRepo.List(&repo.FilterSpec{
		UserId: &userId,
	})

	if err != nil {
		return nil, c.handleError(err)
	}

	var cardsResponse []schema.CardResponse
	for _, card := range cards {
		cardsResponse = append(cardsResponse, *c.cardToResponse(&card))
	}

	return cardsResponse, nil
}

func (c *cardService) Create(userId uuid.UUID, newCard *schema.NewCardRequest) (*schema.CardResponse, error) {
	message := crypto.NewMessage(newCard.Data, "")
	encryptedData, err := message.Encrypt(newCard.Key)
	if err != nil {
		return nil, http.EncryptionError(err)
	}

	hash := sha512.New()
	encryptedKey, err := rsa.EncryptOAEP(hash, rand.Reader, &c.privateKey.PublicKey, []byte(newCard.Key), nil)
	if err != nil {
		return nil, http.EncryptionError(err)
	}

	card, err := c.cardsRepo.Create(&entities.NewCard{
		UserId: userId,
		Title:  newCard.Title,
		Data:   base64.StdEncoding.EncodeToString([]byte(encryptedData)),
		Key:    base64.StdEncoding.EncodeToString(encryptedKey),
	})

	if err != nil {
		return nil, http.InternalError(err)
	}

	return c.cardToResponse(card), nil
}

func (c *cardService) Update(cardId uuid.UUID, updateRequest *schema.UpdateCardRequest) error {
	_, err := c.cardsRepo.Update(cardId, updateRequest.Title)
	if err != nil {
		return c.handleError(err)
	}

	return nil
}

func (c *cardService) Delete(cardId uuid.UUID) error {
	err := c.cardsRepo.Delete(cardId)
	if err != nil {
		return c.handleError(err)
	}

	return nil
}

func (c *cardService) Decrypt(cardId uuid.UUID) (*schema.PlainCardResponse, error) {
	card, err := c.cardsRepo.Get(cardId)
	if err != nil {
		return nil, c.handleError(err)
	}

	hash := sha512.New()

	decodedKey, err := base64.StdEncoding.DecodeString(card.EncryptedKey)
	if err != nil {
		return nil, http.DecodingError(err)
	}

	passphrase, err := rsa.DecryptOAEP(hash, rand.Reader, c.privateKey, decodedKey, nil)
	if err != nil {
		return nil, http.DecryptionError(err)
	}

	decodedData, err := base64.StdEncoding.DecodeString(card.EncryptedData)
	if err != nil {
		return nil, http.DecodingError(err)
	}

	message := crypto.NewMessage("", string(decodedData))
	data, err := message.Decrypt(string(passphrase))
	if err != nil {
		return nil, http.DecryptionError(err)
	}

	return &schema.PlainCardResponse{
		Data: data,
		Key:  string(passphrase),
	}, nil
}

func (c *cardService) ClaimExists(cardId uuid.UUID, userId uuid.UUID) bool {
	return c.cardsRepo.ClaimExists(cardId, userId)
}

func (c *cardService) cardToResponse(card *entities.Card) *schema.CardResponse {
	return &schema.CardResponse{
		ID:            card.ID,
		UserId:        card.UserId,
		EncryptedData: card.EncryptedData,
		CreatedAt:     card.CreatedAt,
		UpdatedAt:     card.UpdatedAt,
	}
}

func (c *cardService) handleError(err error) error {
	if err == sql.ErrNoRows {
		return http.NotFoundError("Card not found")
	} else {
		return http.InternalError(err)
	}
}
