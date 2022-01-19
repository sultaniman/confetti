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

type CardResponse struct {
	ID            uuid.UUID
	UserId        uuid.UUID
	Title         string
	EncryptedData string
	EncryptedKey  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
