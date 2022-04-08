package server

import (
	"time"

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
	}
}
