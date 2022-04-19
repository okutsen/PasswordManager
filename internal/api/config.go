package api

import (
	"fmt"

	"github.com/okutsen/PasswordManager/config"
)

type Config struct {
	Host string
	Port uint
	Addr string
}

func NewConfig(cfg *config.Config) *Config {
	return &Config{
		Host: cfg.Host,
		Port: cfg.Port,
		Addr: fmt.Sprint(cfg.Host, ":", cfg.Port),
	}
}
