package repo

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sultaniman/confetti/platform/entities"
	"time"
)

//go:generate mockgen -source=events.go -destination=../mocks/events.go -package=mocks
type EventRepo interface {
	Add(event *entities.Event)
}

type eventRepo struct {
	Base *Repo
}

const DefaultEventExpirationDuration = time.Hour * 24 * 7

func (e *eventRepo) Add(event *entities.Event) {
	expiresAt := time.Now().UTC().Add(DefaultEventExpirationDuration)
	if event.ExpiresAt != nil {
		expiresAt = *event.ExpiresAt
	}

	query, args, err := e.Base.
		Insert(
			"events",
			"ref",
			"data",
			"expires_at",
			"created_at",
		).
		Values(
			event.Ref,
			event.Data,
			expiresAt,
			time.Now().UTC(),
		).
		ToSql()

	if err != nil {
		log.Err(err).
			Msg(fmt.Sprintf("Unable to build create event query"))
	}

	eventRow := new(entities.Event)
	err = e.Base.DB.Get(eventRow, query, args...)
	if err != nil {
		log.Err(err).
			Str("ref", event.Ref).
			Msg(fmt.Sprintf("Unable to save event"))
	}
}

func NewEventRepo(base *Repo) EventRepo {
	return &eventRepo{Base: base}
}
