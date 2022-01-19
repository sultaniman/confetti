package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/imanhodjaev/confetti/platform/entities"
	"time"
)

//go:generate mockgen -source=cards.go -destination=../mocks/cards.go -package=mocks
type CardRepo interface {
	Get(id uuid.UUID) (*entities.Card, error)
	Create(card *entities.NewCard) (*entities.Card, error)
	Delete(id uuid.UUID) error
}

type cardRepo struct {
	Base *Repo
}

func NewCardRepo(base *Repo) CardRepo {
	return &cardRepo{
		Base: base,
	}
}

func (c *cardRepo) Get(id uuid.UUID) (*entities.Card, error) {
	query, args, err := c.Base.
		Select("cards").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var card entities.Card
	return &card, c.Base.DB.Get(&card, query, args...)
}

func (c *cardRepo) Create(card *entities.NewCard) (*entities.Card, error) {
	query, args, err := c.Base.
		Insert(
			"cards",
			"user_id",
			"title",
			"encrypted_data",
			"encrypted_key",
			"created_at",
			"updated_at",
		).
		Values(
			card.UserId,
			card.Title,
			card.Data,
			card.Key,
			time.Now().UTC(),
			time.Now().UTC(),
		).
		ToSql()

	if err != nil {
		return nil, err
	}

	cardRow := new(entities.Card)
	return cardRow, c.Base.DB.Get(cardRow, query, args...)
}

func (c *cardRepo) Delete(id uuid.UUID) error {
	query, args, err := c.Base.
		Delete("cards", sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	var card *entities.Card
	return c.Base.DB.Get(card, query, args...)
}
