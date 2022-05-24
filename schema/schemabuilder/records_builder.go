package schemabuilder

import (
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/controllerSchema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func BuildControllerRecordFromDBRecord(record *dbschema.Record) controllerSchema.Record {
	return controllerSchema.Record{
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

func BuildControllerRecordsFromDBRecords(records []dbschema.Record) []controllerSchema.Record {
	recordsController := make([]controllerSchema.Record, len(records))
	for i, v := range records {
		recordsController[i] = BuildControllerRecordFromDBRecord(&v)
	}

	return recordsController
}

func BuildAPIRecordFromControllerRecord(record *controllerSchema.Record) apischema.Record {
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

func BuildAPIRecordsFromControllerRecords(records []controllerSchema.Record) []apischema.Record {
	recordsController := make([]apischema.Record, len(records))
	for i, v := range records {
		recordsController[i] = BuildAPIRecordFromControllerRecord(&v)
	}

	return recordsController
}

func BuildControllerRecordFromAPIRecord(record *apischema.Record) controllerSchema.Record {
	return controllerSchema.Record{
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

func BuildDBRecordFromControllerRecord(record *controllerSchema.Record) dbschema.Record {
	return dbschema.Record{
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
