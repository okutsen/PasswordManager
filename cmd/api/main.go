package main

import (
	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/log"
)

// TODO: password tips or reset questions

func main() {
	var logger log.Logger = log.NewLogrusLogger()
	config, err := api.NewConfig()
	if err != nil {
		// TODO: Use default values to configure api
		logger.Fatalf("Failed to initialize config: %s", err.Error())
	}
	serviceAPI := api.New(config, logger)
	err = serviceAPI.Start()
	// close op objects
	logger.Fatal(err)
}
