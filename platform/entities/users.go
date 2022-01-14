package entities

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type NewUser struct {
	FullName string
	Username string
	Email    string
	Password string
	IsAdmin  bool
	IsActive bool
	Provider string
	Settings json.RawMessage
}

type UpdateUser struct {
	FullName string
	Username string
	Email    string
	Settings json.RawMessage
}

type User struct {
	ID        uuid.UUID       `pg:"id"`
	FullName  string          `pg:"full_name"`
	Username  string          `pg:"username"`
	Email     string          `pg:"email"`
	Password  string          `pg:"password"`
	IsAdmin   bool            `pg:"is_admin"`
	IsActive  bool            `pg:"is_active"`
	Settings  json.RawMessage `pg:"settings"`
	Provider  string          `pg:"provider"`
	CreatedAt time.Time       `pg:"created_at"`
	UpdatedAt time.Time       `pg:"updated_at"`
}
