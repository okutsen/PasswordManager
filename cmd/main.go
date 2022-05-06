package main

import (
	"github.com/okutsen/PasswordManager/config"
	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/controller"
	"github.com/okutsen/PasswordManager/internal/log"
)

// TODO: password tips or reset questions

func main() {
	logger := log.NewLogrusLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("initialize config: %v", err)
	}

	ctrl := controller.New(logger)

	serviceAPI := api.New(&api.Config{
		Port: cfg.Port,
	}, ctrl, logger)
	err = serviceAPI.Start()
	if err != nil {
		logger.Fatalf("start application", err)
	}
}
