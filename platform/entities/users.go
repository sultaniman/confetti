package entities

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type NewUser struct {
	FullName string
	Email    string
	Password string
	IsAdmin  bool
	IsActive bool
	Provider string
	Settings json.RawMessage
}

type UpdateUser struct {
	FullName string
	Email    string
	Settings json.RawMessage
}

type User struct {
	ID        uuid.UUID       `db:"id"`
	FullName  string          `db:"full_name"`
	Email     string          `db:"email"`
	Password  string          `db:"password"`
	IsAdmin   bool            `db:"is_admin"`
	IsActive  bool            `db:"is_active"`
	Settings  json.RawMessage `db:"settings"`
	Provider  string          `db:"provider"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
}
