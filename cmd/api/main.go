package main

import (
	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/sirupsen/logrus"
)

// TODO: password tips or reset questions

func main() {
	var config *api.Config = api.NewConfig()
	var log log.Logger = logrus.New() 
	api := api.New(config, log)
	api.Start()
}
