package main

import (
	"github.com/okutsen/PasswordManager/config"
	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/log"
)

// TODO: password tips or reset questions

func main() {
	logger := log.NewLogrusLogger()
	_, err := config.NewConfig("config/config.yaml")
	if err != nil {
		// TODO: Use default values to configure api
		logger.Fatalf("initialize config", err)
	}
	apiConfig, err := api.NewConfig("config/config.yaml")

	serviceAPI := api.New(apiConfig, logger)
	err = serviceAPI.Start()
	// close op objects
	logger.Fatal(err)
}
