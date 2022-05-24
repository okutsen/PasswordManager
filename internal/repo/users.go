package repo

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func (r *Repo) AllUsers() ([]dbschema.User, error) {
	var user []dbschema.User
	result := r.db.Find(&user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all user from db: %w", err)
	}

	return user, nil
}
func (r *Repo) UserByID(id uuid.UUID) (*dbschema.User, error) {
	var user dbschema.User
	result := r.db.First(&user, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user from db: %w", err)
	}

	return &user, nil
}

func (r *Repo) CreateUser(user *dbschema.User) (*dbschema.User, error) {
	user.ID = uuid.New()
	result := r.db.Create(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user in db: %w", err)
	}

	return user, nil
}

func (r *Repo) UpdateUser(user *dbschema.User) (*dbschema.User, error) {
	result := r.db.Model(user).Clauses(clause.Returning{}).Updates(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user in db: %w", err)
	}

	return user, nil
}

func (r *Repo) DeleteUser(id uuid.UUID) (*dbschema.User, error) {
	var user dbschema.User
	result := r.db.Model(&user).Clauses(clause.Returning{}).Delete(&user, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to remove user from db: %w", err)
	}
	return &user, nil
}
