package api

import (
	"fmt"
)

type Config struct {
	Host string
	Port uint
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
