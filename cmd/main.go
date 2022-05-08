package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	serviceAPI := api.New(&api.Config{Port: cfg.Port}, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err = serviceAPI.Start(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("start application %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	closeWaitGroup := &sync.WaitGroup{}

	closeWaitGroup.Add(1)
	err = serviceAPI.Stop(ctx, closeWaitGroup)
	if err != nil {
		panic(err)
	}

	closeWaitGroup.Wait()
}
