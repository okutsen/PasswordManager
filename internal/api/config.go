package api

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port uint
	Addr string
}

func initConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	return viper.ReadInConfig()
}

func NewConfig(configPath string) (*Config, error) {
	if err := initConfig(configPath); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	return &Config{
		Host: viper.GetString("host"),
		Port: viper.GetUint("port"),
	}, nil
}

func (c Config) Address() string {
	c.Addr = fmt.Sprintf("%s:%v", c.Host, c.Port)
	return c.Addr
}
