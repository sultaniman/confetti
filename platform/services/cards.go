package services

import (
	"fmt"
	"github.com/imanhodjaev/confetti/platform/http"
	"github.com/imanhodjaev/confetti/platform/repo"
	"github.com/imanhodjaev/confetti/platform/schema"
	"github.com/imanhodjaev/pwc/crypto"
	"github.com/imanhodjaev/pwc/gen"
)

type CardService interface {
	GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error)
	Save(newCard *schema.NewCardRequest) (*schema.CardResponse, error)
}

type cardService struct {
	usersRepo repo.UserRepo
}

func NewCardService(usersRepo repo.UserRepo) CardService {
	return &cardService{
		usersRepo: usersRepo,
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
	// TODO encrypt card contents with it's key and encrypt key with rsa key and save
	message := crypto.NewMessage(newCard.Data, "")
	encrypted, err := message.Encrypt(newCard.Key)
	if err != nil {
		// TODO: maybe custom error
		return nil, http.InternalError(err)
	}

	fmt.Println(encrypted)
	return nil, err
}
