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
		logger.Fatalf("initialize config: %v", err)
	}

	ctrl := controller.New(logger)

	serviceAPI := api.New(&api.Config{Port: cfg.Port}, ctrl, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = serviceAPI.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("start application %v", err)
			return
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	err = serviceAPI.Stop(ctx)

	logger.Errorf("stop application %v", err)
}
