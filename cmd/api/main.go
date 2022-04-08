package main

import (
	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/sirupsen/logrus"
)

// TODO: password tips or reset questions

func main() {
	var log log.Logger = logrus.New() 
	api := api.NewAPI(log)
	api.Start()
}
