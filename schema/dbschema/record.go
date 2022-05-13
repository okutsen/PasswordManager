package dbschema

import "time"

type User struct {
	ID        uint64    `db:"id"`
	Email     string    `db:"email"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Record struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	Login       string    `db:"login"`
	Password    string    `db:"password"`
	URL         string    `db:"url"`
	Description string    `db:"description"`
	UpdatedBy   string    `db:"updated_by"`
	CreatedBy   string    `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
