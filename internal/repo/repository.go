package repo

import (
	"database/sql"
	"fmt"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

type Repo struct {
	db *sql.DB
}

func New(cfg *Config, logger log.Logger) (*Repo, error) {
	logger.Infof("DB is starting")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("failed to open sql connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to verify connection to DB: %w", err)
	}

	return &Repo{
		db: db}, err
}

func (r *Repo) Close() error {
	return r.db.Close()
}

func AllRecords() ([]dbschema.Record, error) {
	return []dbschema.Record{}, nil
}
