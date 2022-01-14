package schema

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type NewUserRequest struct {
	FullName string
	Username string
	Email    string
	Password string
	Provider string
	Settings json.RawMessage `swaggertype:"object"`
}

type UpdateUserRequest struct {
	FullName string
	Username string
	Settings json.RawMessage `swaggertype:"object"`
}

type UpdateUserEmailRequest struct {
	Email    string
	Password string
}

type UpdateUserPasswordRequest struct {
	OldPassword string
	NewPassword string
}

type UserResponse struct {
	ID        uuid.UUID
	FullName  string
	Username  string
	Email     string
	IsAdmin   bool
	IsActive  bool
	Provider  string
	Settings  json.RawMessage `swaggertype:"object"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersResponse struct {
	Count int
	Users []*UserResponse
}
