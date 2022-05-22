package apischema

import "time"

const (
	InvalidIDParam       = "Invalid ID"
	InvalidJSONMessage   = "Invalid JSON"
	InternalErrorMessage = "Oops, something went wrong"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TODO: add validator
// TODO: add uuid (request id)
type Record struct {
	ID          uint64    `json:"id"`
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
