package schemabuilder

import (
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func BuildAPIRecordFromDBRecord(record *dbschema.Record) apischema.Record {
	return apischema.Record{
		ID:          record.ID,
		Name:        record.Name,
		Login:       record.Login,
		Password:    record.Password,
		URL:         record.URL,
		Description: record.Description,
		CreatedBy:   record.CreatedBy,
		UpdatedBy:   record.UpdatedBy,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func BuildAPIRecordsFromDBRecords(records []dbschema.Record) []apischema.Record {
	recordsAPI := make([]apischema.Record, len(records))
	for i, v := range records {
		recordsAPI[i] = BuildAPIRecordFromDBRecord(&v)
	}

	return recordsAPI
}

func BuildDBRecordFromAPIRecord(record *apischema.Record) dbschema.Record {
	return dbschema.Record{
		ID:          record.ID,
		Name:        record.Name,
		Login:       record.Login,
		Password:    record.Password,
		URL:         record.URL,
		Description: record.Description,
		CreatedBy:   record.CreatedBy,
		UpdatedBy:   record.UpdatedBy,
	}
}
