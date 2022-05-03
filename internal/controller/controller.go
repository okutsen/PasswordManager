package controller

import (
	"fmt"

	"github.com/okutsen/PasswordManager/internal/log"
)

type Controller struct {
	// repo Repo
	log  log.Logger
}

func New(logger log.Logger) *Controller{
	return &Controller{
		log: logger.WithFields(log.Fields{"service": "Controller"}),
	}
}

func (c *Controller) GetAllRecords() (string, error) {
	// queryResult := repo.GetAllRecords()
	queryResult := "Records:\n0,1,2,3,4,5"
	return queryResult, nil
}

func (c *Controller) GetRecord(id uint64) (string, error) {
	// TODO: pass uuid
	// if id < 5 {
	// 	return "", fmt.Errorf("no record with id %d", id)
	// }
	responseBody := fmt.Sprintf("Records:\n%d", id)
	return responseBody, nil
}

func (c *Controller) CreateRecords() (string, error) {
	return "", nil
}
