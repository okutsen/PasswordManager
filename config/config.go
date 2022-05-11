package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port               uint          `envConfig:"PM_PORT" default:"10000"`
	APIShutdownTimeout time.Duration `envConfig:"PM_API_SHUTDOWN_TIMEOUT" default:"30s"`
}

func NewConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("pm", &c)
	if err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	return &c, nil
}
