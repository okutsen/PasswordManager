package main

import (
	"github.com/okutsen/PasswordManager/config"
	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/log"
)

// TODO: password tips or reset questions

func main() {
	logger := log.NewLogrusLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("initialize config: %v", err)
	}

	serviceAPI := api.New(&api.Config{
		Port: cfg.Port,
	}, logger)

	err = serviceAPI.Start()
	if err != nil {
		logger.Fatal(err)
	}
}
