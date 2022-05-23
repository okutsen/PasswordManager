package repo

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

type Data struct {
	repo *Repo
}

func NewRepo(repo *Repo) *Data {
	return &Data{repo: repo}
}

func (r *Data) AllRecords() ([]dbschema.Record, error) {
	var records []dbschema.Record
	result := r.repo.db.Find(&records)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all records from db: %w", err)
	}

	return records, err
}

func (r *Data) CreateRecord(record *dbschema.Record) (*dbschema.Record, error) {
	record.ID = uuid.New()
	result := r.repo.db.Create(record)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create record in db: %w", err)
	}

	return record, err
}

func (r *Data) RecordByID(id uuid.UUID) (*dbschema.Record, error) {
	var record dbschema.Record
	result := r.repo.db.First(&record, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get record from db: %w", err)
	}

	return &record, err
}

func (r *Data) UpdateRecord(record *dbschema.Record) (*dbschema.Record, error) {
	result := r.repo.db.Model(record).Clauses(clause.Returning{}).Updates(record)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update record in db: %w", err)
	}

	return record, err
}

func (r *Data) DeleteRecord(id uuid.UUID) (*dbschema.Record, error) {
	var record dbschema.Record
	result := r.repo.db.Model(&record).Clauses(clause.Returning{}).Delete(&record, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to remove record from db: %w", err)
	}

	return &record, err
}

func (r *Data) AllUsers() ([]dbschema.User, error) {
	var user []dbschema.User
	result := r.repo.db.Find(&user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all user from db: %w", err)
	}

	return user, err
}

func (r *Data) CreateUser(user *dbschema.User) (*dbschema.User, error) {
	user.ID = uuid.New()
	result := r.repo.db.Create(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user in db: %w", err)
	}

	return user, err
}

func (r *Data) UserByID(id uuid.UUID) (*dbschema.User, error) {
	var user dbschema.User
	result := r.repo.db.First(&user, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user from db: %w", err)
	}

	return &user, err
}

func (r *Data) UpdateUser(user *dbschema.User) (*dbschema.User, error) {
	result := r.repo.db.Model(user).Clauses(clause.Returning{}).Updates(user)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update user in db: %w", err)
	}

	return user, err
}

func (r *Data) DeleteUser(id uuid.UUID) (*dbschema.User, error) {
	var user dbschema.User
	result := r.repo.db.Model(&user).Clauses(clause.Returning{}).Delete(&user, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to remove user from db: %w", err)
	}
	return &user, err
}
