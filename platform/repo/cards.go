package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/sultaniman/confetti/platform/entities"
	"time"
)

type FilterSpec struct {
	UserId *uuid.UUID
	ID     *uuid.UUID
}

//go:generate mockgen -source=cards.go -destination=../mocks/cards.go -package=mocks
type CardRepo interface {
	Get(id uuid.UUID) (*entities.Card, error)
	List(filterSpec *FilterSpec) ([]entities.Card, error)
	Create(card *entities.NewCard) (*entities.Card, error)
	Update(cardId uuid.UUID, newTitle string) (*entities.Card, error)
	Delete(id uuid.UUID) error
	ClaimExists(cardId uuid.UUID, userId uuid.UUID) bool
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

	card := new(entities.Card)
	return card, c.Base.DB.Get(card, query, args...)
}

func (c *cardRepo) List(filterSpec *FilterSpec) ([]entities.Card, error) {
	qs := c.Base.
		Select("cards")

	filters := sq.Eq{}
	if filterSpec.ID != nil {
		filters["id"] = filterSpec.ID
	}

	if filterSpec.UserId != nil {
		filters["user_id"] = filterSpec.UserId
	}

	query, args, err := qs.
		Where(filters).
		ToSql()

	if err != nil {
		return nil, err
	}

	cards := new([]entities.Card)
	return *cards, c.Base.DB.Select(cards, query, args...)
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

func (c *cardRepo) Update(cardId uuid.UUID, newTitle string) (*entities.Card, error) {
	query, args, err := c.Base.
		Update("cards", true).
		Where(sq.Eq{"id": cardId}).
		Set("title", newTitle).
		ToSql()

	if err != nil {
		return nil, err
	}

	card := new(entities.Card)
	return card, c.Base.DB.Get(card, query, args...)
}

func (c *cardRepo) Delete(id uuid.UUID) error {
	query, args, err := c.Base.
		Delete("cards", sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	card := new(entities.Card)
	return c.Base.DB.Get(card, query, args...)
}

func (c *cardRepo) ClaimExists(cardId uuid.UUID, userId uuid.UUID) bool {
	query, args, err := c.Base.
		Count("cards", sq.Eq{"id": cardId, "user_id": userId}).
		ToSql()

	if err != nil {
		return false
	}

	rowCount := 0
	err = c.Base.DB.Get(&rowCount, query, args...)
	if err != nil {
		return false
	}

	return rowCount > 0
}
