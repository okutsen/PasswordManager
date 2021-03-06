package repo

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func (r *Repo) AllRecords() ([]dbschema.Record, error) {
	var records []dbschema.Record
	result := r.db.Find(&records)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all records from db: %w", err)
	}

	return records, nil
}
func (r *Repo) RecordByID(id uuid.UUID) (*dbschema.Record, error) {
	var record dbschema.Record
	result := r.db.First(&record, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to get record from db: %w", err)
	}

	return &record, nil
}

func (r *Repo) CreateRecord(record *dbschema.Record) (*dbschema.Record, error) {
	record.ID = uuid.New()
	result := r.db.Create(record)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to create record in db: %w", err)
	}

	return record, nil
}

func (r *Repo) UpdateRecord(record *dbschema.Record) (*dbschema.Record, error) {
	result := r.db.Model(record).Clauses(clause.Returning{}).Updates(record)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to update record in db: %w", err)
	}

	return record, nil
}

func (r *Repo) DeleteRecord(id uuid.UUID) (*dbschema.Record, error) {
	var record dbschema.Record
	result := r.db.Model(&record).Clauses(clause.Returning{}).Delete(&record, id)
	err := result.Error
	if err != nil {
		return nil, fmt.Errorf("failed to remove record from db: %w", err)
	}

	return &record, nil
}
