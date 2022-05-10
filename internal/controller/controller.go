package controller

import (
	"github.com/okutsen/PasswordManager/schema/dbschema"
	"github.com/okutsen/PasswordManager/internal/log"
)

type Controller struct {
	// repo Repo
	log log.Logger
}

func New(logger log.Logger) *Controller {
	return &Controller{
		log: logger.WithFields(log.Fields{"service": "Controller"}),
	}
}

func (c *Controller) GetAllRecords() ([]dbschema.Record, error) {
	// queryResult := repo.GetAllRecords()
	return []dbschema.Record{{
		ID:       1,
		Name:     "testName",
		Login:    "testLogin",
		Password: "testPassword",
	},
		{
			ID:       2,
			Name:     "testName",
			Login:    "testLogin",
			Password: "testPassword",
		},
		{
			ID:       3,
			Name:     "testName",
			Login:    "testLogin",
			Password: "testPassword",
		},
		{
			ID:       4,
			Name:     "testName",
			Login:    "testLogin",
			Password: "testPassword",
		},
		{
			ID:       5,
			Name:     "testName",
			Login:    "testLogin",
			Password: "testPassword",
		}}, nil
}

func (c *Controller) GetRecord(id uint64) ([]dbschema.Record, error) {
	// TODO: pass uuid
	return []dbschema.Record{{
		ID:       id,
		Name:     "testName",
		Login:    "testLogin",
		Password: "testPassword",
	},}, nil
}

func (c *Controller) CreateRecords(records []dbschema.Record) error {
	return nil
}
