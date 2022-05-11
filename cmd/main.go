package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
		logger.Errorf("failed to initialize config: %v", err)
		os.Exit(1)
	}

	ctrl := controller.New(logger)

	serviceAPI := api.New(&api.Config{Port: cfg.Port}, ctrl, logger)

	go func() {
		err = serviceAPI.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("failed to start application %v", err)
			return
		}
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	osCall := <-osSignals
	logger.Infof("system call: %v", osCall)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.APIShutdownTimeout)
	defer cancel()

	err = serviceAPI.Stop(ctx)
	if err != nil {
		logger.Errorf("failed to stop application %v", err)
		os.Exit(1)
	}

}
