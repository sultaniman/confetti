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
	Data      string
	Key       string
	ExpiresIn int64
	Image     []byte
}

type NewCardRequest struct {
	Data      string
	Key       string
	ExpiresIn int64
}

type CardResponse struct {
	ID            uuid.UUID
	UserId        uuid.UUID
	EncryptedData string
	EncryptedKey  string
	ExpiresIn     int64
	CreatedAt     time.Time
}
