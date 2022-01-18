package entities

import (
	"github.com/google/uuid"
	"time"
)

type NewCard struct {
	UserId    uuid.UUID
	Data      string
	Key       string
	ExpiresIn int64
}

type ExpirationUpdate struct {
	ExpiresIn int64
}

type Card struct {
	ID            uuid.UUID `db:"id"`
	UserId        uuid.UUID `db:"user_id"`
	EncryptedData string    `db:"encrypted_data"`
	EncryptedKey  string    `db:"encrypted_key"`
	ExpiresIn     int64     `db:"expires_in"`
	CreatedAt     time.Time `db:"created_at"`
}
