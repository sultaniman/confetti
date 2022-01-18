package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"github.com/google/uuid"
	"github.com/imanhodjaev/confetti/platform/http"
	"github.com/imanhodjaev/confetti/platform/repo"
	"github.com/imanhodjaev/confetti/platform/schema"
	"github.com/imanhodjaev/pwc/crypto"
	"github.com/imanhodjaev/pwc/gen"
	"time"
)

type CardService interface {
	GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error)
	Save(newCard *schema.NewCardRequest) (*schema.CardResponse, error)
}

type cardService struct {
	privateKey *rsa.PrivateKey
	usersRepo  repo.UserRepo
}

func NewCardService(usersRepo repo.UserRepo, privateKey *rsa.PrivateKey) CardService {
	return &cardService{
		privateKey: privateKey,
		usersRepo:  usersRepo,
	}
}

func (c *cardService) GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error) {
	_, card, err := gen.GenerateClassicCard(options.IncludeSymbols, options.DigitsArea, false)
	if err != nil {
		return nil, http.InternalError(err)
	}

	return &schema.NewCardResponse{
		Data:      string(card.GetBytes()),
		Key:       card.Passphrase,
		ExpiresIn: 0,
	}, nil
}

func (c *cardService) Save(newCard *schema.NewCardRequest) (*schema.CardResponse, error) {
	message := crypto.NewMessage(newCard.Data, "")
	encryptedData, err := message.Encrypt(newCard.Key)
	if err != nil {
		// TODO: maybe custom error
		return nil, http.InternalError(err)
	}

	hash := sha512.New()
	encryptedKey, err := rsa.EncryptOAEP(hash, rand.Reader, &c.privateKey.PublicKey, []byte(newCard.Key), nil)
	if err != nil {
		// TODO: maybe custom error
		return nil, http.InternalError(err)
	}

	// TODO: save
	return &schema.CardResponse{
		UserId:        uuid.New(),
		EncryptedData: encryptedData,
		EncryptedKey:  string(encryptedKey),
		ExpiresIn:     newCard.ExpiresIn,
		CreatedAt:     time.Now(),
	}, err
}
