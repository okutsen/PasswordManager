package main

import (
	"github.com/sirupsen/logrus"

	"github.com/okutsen/PasswordManager/internal/api"
	"github.com/okutsen/PasswordManager/internal/log"
)

// TODO: password tips or reset questions

func main() {
	var logger log.Logger = logrus.New()
	config, err := api.NewConfig()
	if err != nil {
		// TODO: Use default values to configure api
		logger.Fatalf("failed to initialize config: %v", err)
	}
	serviceAPI := api.New(config, logger)
	err = serviceAPI.Start()
	// close op objects
	logger.Fatal(err)
}
