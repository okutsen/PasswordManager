package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port uint
}

func NewConfig() (*Config, error) {
	err := godotenv.Load("config/.env")
	viper.SetEnvPrefix("pm")
	viper.AutomaticEnv()
	return &Config{
		Host: viper.GetString("host"),
		Port: viper.GetUint("port"),
	}, err
}
