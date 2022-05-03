package controller

import (
	"fmt"

	"github.com/okutsen/PasswordManager/internal/log"
)

type ControllerService struct {
	// repo Repo
	log  log.Logger
}

func New(logger log.Logger) *ControllerService{
	return &ControllerService{
		log: logger.WithFields(log.Fields{"service": "Controller"}),
	}
}

func (c *ControllerService) GetAllRecords() (string, error) {
	// queryResult := repo.GetAllRecords()
	queryResult := "Records:\n0,1,2,3,4,5"
	return queryResult, nil
}

func (c *ControllerService) GetRecord(id uint) (string, error) {
	// TODO: pass uuid
	responseBody := fmt.Sprintf("Records:\n%d", id)
	return responseBody, nil
}

func (c *ControllerService) CreateRecords() (string, error) {
	return "", nil
}
