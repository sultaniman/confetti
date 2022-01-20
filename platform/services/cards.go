package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/imanhodjaev/confetti/platform/entities"
	"github.com/imanhodjaev/confetti/platform/http"
	"github.com/imanhodjaev/confetti/platform/repo"
	"github.com/imanhodjaev/confetti/platform/schema"
	"github.com/imanhodjaev/pwc/crypto"
	"github.com/imanhodjaev/pwc/gen"
)

type CardService interface {
	Generate(options *schema.CardOptions) (*schema.NewCardResponse, error)
	Get(cardId uuid.UUID) (*schema.CardResponse, error)
	Create(userId uuid.UUID, newCard *schema.NewCardRequest) (*schema.CardResponse, error)
	Update(cardId uuid.UUID, updateRequest *schema.UpdateCardRequest) error
	Delete(cardId uuid.UUID) error
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

func (c *cardService) cardToResponse(card *entities.Card) *schema.CardResponse {
	return &schema.CardResponse{
		ID:            card.ID,
		UserId:        card.UserId,
		EncryptedData: card.EncryptedData,
		EncryptedKey:  card.EncryptedKey,
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
