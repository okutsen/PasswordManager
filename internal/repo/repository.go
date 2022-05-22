package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
