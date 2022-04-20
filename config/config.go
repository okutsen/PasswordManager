package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port uint
}

func initConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	return viper.ReadInConfig()
}

func NewConfig(configPath string) (*Config, error) {
	if err := initConfig(configPath); err != nil {
		return nil, fmt.Errorf("failed to init config: %w", err)
	}
	return &Config{
		Host: viper.GetString("host"),
		Port: viper.GetUint("port"),
	}, nil
}
