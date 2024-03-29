package schema

import (
	"github.com/google/uuid"
	"time"
)

type CardOptions struct {
	IncludeSymbols bool
	DigitsArea     bool
}

type NewCardResponse struct {
	Data string
	Key  string
}

type NewCardRequest struct {
	Title string
	Data  string
	Key   string
}

type UpdateCardRequest struct {
	Title string
}

type CardResponse struct {
	ID            uuid.UUID
	UserId        uuid.UUID
	Title         string
	EncryptedData string
	KeyID         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CardClaim struct {
	CardId uuid.UUID
	UserId uuid.UUID
	Exists bool
}

type PlainCardResponse struct {
	Title string
	Data  string
	Key   string
}
