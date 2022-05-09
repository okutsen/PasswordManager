package controller

import (
	"github.com/okutsen/PasswordManager/domain"
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

func (c *Controller) GetAllRecords() ([]domain.Record, error) {
	// queryResult := repo.GetAllRecords()
	return []domain.Record{{
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

func (c *Controller) GetRecord(id uint64) ([]domain.Record, error) {
	// TODO: pass uuid
	return []domain.Record{{
		ID:       id,
		Name:     "testName",
		Login:    "testLogin",
		Password: "testPassword",
	},}, nil
}

func (c *Controller) CreateRecords(records []domain.Record) error {
	return nil
}
