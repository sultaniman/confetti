package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
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
	GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error)
	Create(userId uuid.UUID, newCard *schema.NewCardRequest) (*schema.CardResponse, error)
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

func (c *cardService) GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error) {
	_, card, err := gen.GenerateClassicCard(options.IncludeSymbols, options.DigitsArea, false)
	if err != nil {
		return nil, http.InternalError(err)
	}

	return &schema.NewCardResponse{
		Data: string(card.GetBytes()),
		Key:  card.Passphrase,
	}, nil
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

func (c *cardService) Delete(cardId uuid.UUID) error {
	err := c.cardsRepo.Delete(cardId)
	if err != nil {
		return http.InternalError(err)
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
