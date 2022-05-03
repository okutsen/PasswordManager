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
