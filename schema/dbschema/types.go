package dbschema

import "github.com/google/uuid"

// TODO: more fields (dateCreated, url, description)
type Record struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}
