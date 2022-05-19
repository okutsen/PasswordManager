package repo

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/okutsen/PasswordManager/schema/dbschema"
)

type Repo struct {
	db *gorm.DB
}

func New(cfg *Config) (*Repo, error) {
	db, err := gorm.Open(postgres.Open(cfg.Address()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Repo{
		db: db,
	}, err
}

func (r *Repo) Close() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (r *Repo) AllRecordsFromDB() ([]dbschema.Record, error) {
	var records []dbschema.Record
	result := r.db.Find(&records).Order("id")
	err := result.Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r Repo) CreateRecordInDB() (*dbschema.Record, error) {
	records := &dbschema.Record{
		ID:          0,
		Name:        "John",
		Login:       "John1823",
		Password:    "12345",
		URL:         "john1823.com",
		Description: "My site",
		UpdatedBy:   "User",
		CreatedBy:   "John1823",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	result := r.db.Create(&records)
	err := result.Error()
	if err != nil {
		return nil, err
	}
	return records, err
}
