package repo

import "fmt"

type Config struct {
	Host     string
	Port     string
	DBName   string
	Username string
	SSLMode  string
	Password string
}

func (c *Config) Address() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.Username, c.Password, c.DBName, c.Port, c.SSLMode)
}
