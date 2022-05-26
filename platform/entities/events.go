package entities

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID        *uuid.UUID      `pg:"type:uuid"`
	Ref       string          `pg:"ref"` // i.e. Origin ID
	Data      json.RawMessage `pg:"data"`
	ExpiresAt *time.Time      `pg:"expires_at"`
	CreatedAt *time.Time      `pg:"created_at"`
}
