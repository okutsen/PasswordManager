package controller

import (
	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/dbschema"
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

func (c *Controller) GetAllRecords() ([]*dbschema.Record, error) {
	// queryResult := repo.GetAllRecords()
	return []*dbschema.Record{
		{
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
		},
	}, nil
}

func (c *Controller) GetRecord(id uint64) (*dbschema.Record, error) {
	// TODO: pass uuid
	return &dbschema.Record{
		ID:       id,
		Name:     "testName",
		Login:    "testLogin",
		Password: "testPassword",
	}, nil
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)
func (c *Controller) CreateRecord(records *dbschema.Record) error {
	return nil
}
// 200, 204(if no changes?), 404
func (c *Controller) UpdateRecord(records *dbschema.Record) error {
	return nil
}
// 200, 404
func (c *Controller) DeleteRecord(ids uint64) error {
	return nil
}
