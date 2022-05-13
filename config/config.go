package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port            uint          `envConfig:"PM_PORT" default:"10000"`
	ShutdownTimeout time.Duration `envConfig:"PM_API_SHUTDOWN_TIMEOUT" default:"30s"`
}

type DBConfig struct {
	Host     string `envConfig:"PM_DB_HOST" default:"postgres"`
	Port     string `envConfig:"PM_DB_PORT" default:"5432"`
	DBName   string `envConfig:"PM_PM" default:"PM"`
	Username string `envConfig:"PM_USERNAME" default:"admin"`
	SSLMode  string `envConfig:"SSL_MODE" default:"disable"`
	Password string `envConfig:"PM_PASSWORD" default:"1234"`
}

func NewConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("pm", &c)
	if err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	return &c, nil
}
