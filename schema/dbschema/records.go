package dbschema

import "time"

type Users struct {
	ID        uint64    `db:"id"`
	Email     string    `db:"email"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`
}

type Records struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	Login       string    `db:"login"`
	Password    string    `db:"password"`
	URL         string    `db:"url"`
	Description string    `db:"description"`
	UpdatedBy   string    `db:"updatedby"`
	CreatedBy   string    `db:"createdby"`
	CreatedAt   time.Time `db:"createdate"`
	UpdatedAt   time.Time `db:"updatedat"`
}
