package schema

import (
	"encoding/json"
	"github.com/google/uuid"
)

type UserEventData struct {
	ID          uuid.UUID       `json:"ID"`
	FullName    string          `json:"FullName"`
	IsActive    bool            `json:"IsActive"`
	IsConfirmed bool            `json:"IsConfirmed"`
	Provider    string          `json:"Provider"`
	Settings    json.RawMessage `json:"Settings"`
}

type CardEventData struct {
	ID            uuid.UUID `json:"ID"`
	UserId        uuid.UUID `json:"UserId"`
	Title         string    `json:"Title"`
	EncryptedData string    `json:"EncryptedData"`
	KeyID         string    `json:"KeyID"`
}
