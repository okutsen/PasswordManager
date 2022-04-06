package server

import (
	"time"

	"github.com/okutsen/PasswordManager/config"
)

type Config struct {
	Ports
	Timings
}

type Ports struct {
	ServerListenPort string
}
type Timings struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		Ports: Ports{
			ServerListenPort: config.ServerListenPort,
		},
		Timings: Timings{
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
		},
	}
}
