package controller

import (
	"github.com/google/uuid"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

type RecordsRepo interface {
	AllRecords() ([]dbschema.Record, error)
	RecordByID(id uuid.UUID) (*dbschema.Record, error)
	CreateRecord(record *dbschema.Record) (*dbschema.Record, error)
	UpdateRecord(record *dbschema.Record) (*dbschema.Record, error)
	DeleteRecord(id uuid.UUID) (*dbschema.Record, error)
}

type RecordsController struct {
	records RecordsRepo
	log     log.Logger
}

func NewRecords(logger log.Logger, repo RecordsRepo) *RecordsController {
	return &RecordsController{
		log:     logger.WithFields(log.Fields{"service": "Controller"}),
		records: repo,
	}
}

// AllRecords get all records from DB
func (c *RecordsController) AllRecords() ([]apischema.Record, error) {
	getRecords, err := c.records.AllRecords()
	if err != nil {
		return nil, err
	}

	recordsAPI := schemabuilder.BuildAPIRecordsFromDBRecords(getRecords)
	return recordsAPI, err
}

// Record get one record from DB by ID
func (c *RecordsController) Record(id uuid.UUID) (*apischema.Record, error) {
	getRecord, err := c.records.RecordByID(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	recordAPI := schemabuilder.BuildAPIRecordFromDBRecord(getRecord)
	return &recordAPI, err
}

// CreateRecord creates new record in DB
func (c *RecordsController) CreateRecord(record *apischema.Record) (*apischema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromAPIRecord(record)
	createdDBRecord, err := c.records.CreateRecord(&dbRecord)
	if err != nil {
		return nil, err
	}

	createdAPIRecord := schemabuilder.BuildAPIRecordFromDBRecord(createdDBRecord)
	return &createdAPIRecord, err
}

// UpdateRecord updates record in DB
func (c *RecordsController) UpdateRecord(id uuid.UUID, record *apischema.Record) (*apischema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromAPIRecord(record)
	dbRecord.ID = id

	updatedRecord, err := c.records.UpdateRecord(&dbRecord)
	if err != nil {
		return nil, err
	}

	updatedApiRecord := schemabuilder.BuildAPIRecordFromDBRecord(updatedRecord)
	return &updatedApiRecord, err
}

// DeleteRecord deletes record in DB
func (c *RecordsController) DeleteRecord(id uuid.UUID) (*apischema.Record, error) {
	dbRecord, err := c.records.DeleteRecord(id)
	if err != nil {
		return nil, err
	}

	record := schemabuilder.BuildAPIRecordFromDBRecord(dbRecord)

	return &record, err
}
