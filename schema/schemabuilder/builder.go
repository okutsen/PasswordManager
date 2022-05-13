package schemabuilder

import (
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func BuildRecordsAPIFrom(records []dbschema.Record) []apischema.Record {
	recordsAPI := make([]apischema.Record, len(records))
	for i, v := range records {
		recordsAPI[i] = apischema.Record{
			ID:       v.ID,
			Name:     v.Name,
			Login:    v.Login,
			Password: v.Password,
		}
	}
	return recordsAPI
}

func BuildRecordsFrom(recordsAPI []apischema.Record) []dbschema.Record {
	records := make([]dbschema.Record, len(recordsAPI))
	for i, v := range recordsAPI {
		records[i] = dbschema.Record{
			ID:       v.ID,
			Name:     v.Name,
			Login:    v.Login,
			Password: v.Password,
		}
	}
	return records
}
