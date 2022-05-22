package repo

import (
	"fmt"

	"gorm.io/gorm/clause"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

type Records struct {
	repo *Repo
}

func NewRecordsRepo(repo *Repo) *Records {
	return &Records{repo: repo}
}

func (r *Records) AllRecordsFromDB() ([]dbschema.Record, error) {
	var records []dbschema.Record
	result := r.repo.db.Find(&records)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all records from db: %w", err)
	}

	return records, err
}

func (r *Records) CreateRecordInDB(record *dbschema.Record) (*dbschema.Record, error) {
	result := r.repo.db.Create(record)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create record in db: %w", err)
	}

	return record, err
}

func (r *Records) RecordByIDFromDB(id uint64) (*dbschema.Record, error) {
	var record dbschema.Record
	result := r.repo.db.First(&record, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get record from db: %w", err)
	}

	return &record, err
}

func (r *Records) UpdateRecordInDB(record *dbschema.Record) (*dbschema.Record, error) {
	result := r.repo.db.Model(record).Clauses(clause.Returning{}).Updates(record)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update record in db: %w", err)
	}

	return record, err
}

func (r *Records) DeleteRecordFromDB(id uint64) error {
	var record dbschema.Record
	result := r.repo.db.Delete(&record, id)
	err := result.Error
	if err != nil {
		return fmt.Errorf("failed to remove record from db: %w", err)
	}

	return err
}
