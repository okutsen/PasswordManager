package schemabuilder

import (
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func BuildRecordAPIFrom(record *dbschema.Record) apischema.Record {
	return apischema.Record{
		ID:       record.ID,
		Name:     record.Name,
		Login:    record.Login,
		Password: record.Password,
	}
}

func BuildRecordsAPIFrom(records []dbschema.Record) []apischema.Record {
	recordsAPI := make([]apischema.Record, len(records))
	for i, v := range records {
		recordsAPI[i] = BuildRecordAPIFrom(&v)
	}
	return recordsAPI
}

func BuildRecordFrom(recordAPI *apischema.Record) dbschema.Record {
	return dbschema.Record{
		ID:       recordAPI.ID,
		Name:     recordAPI.Name,
		Login:    recordAPI.Login,
		Password: recordAPI.Password,
	}
}

func BuildRecordsFrom(recordsAPI []apischema.Record) []dbschema.Record {
	records := make([]dbschema.Record, len(recordsAPI))
	for i, v := range recordsAPI {
		records[i] = BuildRecordFrom(&v)
	}
	return records
}
