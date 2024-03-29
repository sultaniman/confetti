package schema

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type NewUserRequest struct {
	FullName string
	Email    string
	Password string
	Provider string
	Settings json.RawMessage `swaggertype:"object"`
}

type UpdateUserRequest struct {
	FullName string
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
	ID          uuid.UUID
	FullName    string
	Email       string
	IsAdmin     bool
	IsActive    bool
	IsConfirmed bool
	Provider    string
	Settings    json.RawMessage `swaggertype:"object"`
	Password    string          `json:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UsersResponse struct {
	Count int
	Users []*UserResponse
}

type ActionCodeRequest struct {
	Type  string
	Email string
}

type ActionCode struct {
	ID        uuid.UUID
	UserId    uuid.UUID
	Code      string
	CreatedAt time.Time `json:"-"`
}
