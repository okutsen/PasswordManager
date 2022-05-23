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
	"github.com/okutsen/PasswordManager/internal/repo"
)

// TODO: password tips or reset questions

func main() {
	logger := log.NewLogrusLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Errorf("failed to initialize config: %v", err)
		os.Exit(1)
	}

	db, err := repo.New(&repo.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		DBName:   cfg.DB.DBName,
		Username: cfg.DB.Username,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logger.Errorf("failed to initialize DB: %v", err)
		os.Exit(1)
	}
	logger.Info("DB is started")

	repository := repo.NewRepo(db)

	ctrl := controller.NewController(logger, repository)

	serviceAPI := api.New(&api.Config{Port: cfg.API.Port}, ctrl, logger)

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

	err = db.Close()
	if err != nil {
		logger.Warnf("failed to close DB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.API.ShutdownTimeout)
	defer cancel()

	err = serviceAPI.Stop(ctx)
	if err != nil {
		logger.Errorf("failed to stop application %v", err)
		os.Exit(1)
	}

}
