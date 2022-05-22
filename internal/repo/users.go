package repo

import (
	"fmt"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

type Users struct {
	repo *Repo
}

func NewUsersRepo(repo *Repo) *Users {
	return &Users{repo: repo}
}

func (r *Users) AllUsersFromDB() ([]dbschema.User, error) {
	var user []dbschema.User
	result := r.repo.db.Find(&user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all user from db: %w", err)
	}

	return user, err
}

func (r *Users) CreateUserInDB(user *dbschema.User) (*dbschema.User, error) {
	result := r.repo.db.Create(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user in db: %w", err)
	}

	return user, err
}

func (r *Users) UserFromDB(id uint64) (*dbschema.User, error) {
	var user dbschema.User
	result := r.repo.db.First(&user, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user from db: %w", err)
	}

	return &user, err
}

func (r *Users) UpdateUserInDB(user *dbschema.User) (*dbschema.User, error) {
	result := r.repo.db.Model(user).Updates(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user in db: %w", err)
	}

	return user, err
}

func (r *Users) DeleteUserFromDB(id uint64) error {
	var user dbschema.User
	result := r.repo.db.Delete(&user, id)
	err := result.Error
	if err != nil {
		return fmt.Errorf("failed to remove user from db: %w", err)
	}

	return err
}
