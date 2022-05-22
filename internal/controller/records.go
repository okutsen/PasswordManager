package controller

import (
	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

type RecordsRepo interface {
	AllRecordsFromDB() ([]dbschema.Record, error)
	RecordByIDFromDB(id uint64) (*dbschema.Record, error)
	CreateRecordInDB(record *dbschema.Record) (*dbschema.Record, error)
	UpdateRecordInDB(record *dbschema.Record) (*dbschema.Record, error)
	DeleteRecordFromDB(id uint64) error
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

func (c *RecordsController) AllRecords() ([]apischema.Record, error) {
	getRecords, err := c.records.AllRecordsFromDB()
	if err != nil {
		return nil, err
	}

	recordsAPI := schemabuilder.BuildAPIRecordsFromDBRecords(getRecords)
	return recordsAPI, err
}

func (c *RecordsController) Record(id uint64) (*apischema.Record, error) {
	getRecord, err := c.records.RecordByIDFromDB(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	recordAPI := schemabuilder.BuildAPIRecordFromDBRecord(getRecord)
	return &recordAPI, err
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)
func (c *RecordsController) CreateRecord(record *apischema.Record) (*apischema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromAPIRecord(record)
	createdDBRecord, err := c.records.CreateRecordInDB(&dbRecord)
	if err != nil {
		return nil, err
	}

	createdAPIRecord := schemabuilder.BuildAPIRecordFromDBRecord(createdDBRecord)
	return &createdAPIRecord, err
}

// 200, 204(if no changes?), 404
func (c *RecordsController) UpdateRecord(id uint64, record *apischema.Record) (*apischema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromAPIRecord(record)
	dbRecord.ID = id

	updatedRecord, err := c.records.UpdateRecordInDB(&dbRecord)
	if err != nil {
		return nil, err
	}

	updatedApiRecord := schemabuilder.BuildAPIRecordFromDBRecord(updatedRecord)
	return &updatedApiRecord, err
}

// 200, 404
func (c *RecordsController) DeleteRecord(id uint64) error {
	return c.records.DeleteRecordFromDB(id)
}
