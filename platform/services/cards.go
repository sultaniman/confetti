package services

import (
	"bytes"
	"github.com/imanhodjaev/confetti/platform/http"
	"github.com/imanhodjaev/confetti/platform/repo"
	"github.com/imanhodjaev/confetti/platform/schema"
	"github.com/imanhodjaev/pwc/gen"
	"image/jpeg"
)

type CardService interface {
	GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error)
	Create(newCard *schema.NewCardRequest) (*schema.CardResponse, error)
}

type cardService struct {
	usersRepo repo.UserRepo
}

func NewCardService(usersRepo repo.UserRepo) CardService {
	return &cardService{
		usersRepo: usersRepo,
	}
}

func (c cardService) GenerateCard(options *schema.CardOptions) (*schema.NewCardResponse, error) {
	canvas, card, err := gen.GenerateClassicCard(options.IncludeSymbols, options.DigitsArea)
	if err != nil {
		return nil, http.InternalError(err)
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, canvas.Context.Image(), nil)
	if err != nil {
		return nil, http.InternalError(err)
	}

	return &schema.NewCardResponse{
		Data:      string(card.GetBytes()),
		Key:       card.Passphrase,
		ExpiresIn: 0,
		Image:     buf.Bytes(),
	}, nil
}

func (c cardService) Create(newCard *schema.NewCardRequest) (*schema.CardResponse, error) {
	//TODO implement me
	panic("implement me")
}
