package repo

import (
	"fmt"

	"gorm.io/gorm/clause"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

type Users struct {
	repo *Repo
}

func NewUsersRepo(repo *Repo) *Users {
	return &Users{repo: repo}
}

func (r *Users) AllUsers() ([]dbschema.User, error) {
	var user []dbschema.User
	result := r.repo.db.Find(&user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all user from db: %w", err)
	}

	return user, err
}

func (r *Users) CreateUser(user *dbschema.User) (*dbschema.User, error) {
	result := r.repo.db.Create(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user in db: %w", err)
	}

	return user, err
}

func (r *Users) UserByID(id uint64) (*dbschema.User, error) {
	var user dbschema.User
	result := r.repo.db.First(&user, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user from db: %w", err)
	}

	return &user, err
}

func (r *Users) UpdateUser(user *dbschema.User) (*dbschema.User, error) {
	result := r.repo.db.Model(user).Clauses(clause.Returning{}).Updates(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user in db: %w", err)
	}

	return user, err
}

func (r *Users) DeleteUser(id uint64) error {
	var user dbschema.User
	result := r.repo.db.Delete(&user, id)
	err := result.Error
	if err != nil {
		return fmt.Errorf("failed to remove user from db: %w", err)
	}

	return err
}