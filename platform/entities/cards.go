package entities

import (
	"github.com/google/uuid"
	"time"
)

type NewCard struct {
	UserId uuid.UUID
	Title  string
	Data   string
	Key    string
}

type TitleUpdate struct {
	Title string
}

type Card struct {
	ID            uuid.UUID `db:"id"`
	UserId        uuid.UUID `db:"user_id"`
	Title         string    `db:"title"`
	EncryptedData string    `db:"encrypted_data"`
	EncryptedKey  string    `db:"encrypted_key"`
	CreatedAt     time.Time `db:"created_at"`
}
