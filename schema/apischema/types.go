package apischema

import (
	"time"

	"github.com/google/uuid"
)

const (
	InvalidJSONMessage     = "Invalid JSON"
	InvalidRecordIDMessage = "Invalid record ID"
	InvalidUserIDMessage = "Invalid user ID"
	InternalErrorMessage   = "Oops, something went wrong"
	UnAuthorizedMessage    = "Sign in to use service"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Record struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Error struct {
	Message string `json:"message"`
}
