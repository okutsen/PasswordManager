package api

import (
	"fmt"
)

type Config struct {
	Port uint
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c Config) LocalAddress() string {
	return fmt.Sprintf("127.0.0.1:%d", c.Port)
}
