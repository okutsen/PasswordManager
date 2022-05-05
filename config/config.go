package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port uint `envConfig:"PM_PORT" default:"10000"`
}

func NewConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("pm", &c)
	if err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	return &c, nil
}
