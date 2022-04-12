package api

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func NewConfig() (*Config, error) {
	if err := initConfig(); err != nil {
		return nil, fmt.Errorf("failed to init config: %w", err)
	}
	return &Config{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}, nil
}
