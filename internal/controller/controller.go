package controller

import "github.com/okutsen/PasswordManager/internal/log"

type Repo interface {
	GetAllRecords()
	GetRecord()
	CreateRecords()
}

type Controller struct {
	repo Repo
	log  log.Logger
}

func (c *Controller) GetAllRecords() string {
	// repo.GetAllRecords()
	queryResult := "Records:\n0,1,2,3,4,5"
	return queryResult
}

func (c *Controller) GetRecord(id string) string {
	// 
	queryResult := "Records:\n0,1,2,3,4,5"
	return queryResult
}

func (c *Controller) CreateRecords() string {
	queryResult := "Records:\n0,1,2,3,4,5"
	return queryResult
}
