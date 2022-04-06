package api

import (
	"time"

	"github.com/okutsen/PasswordManager/config"
)

type Config struct {
	Ports
	PathVars
	ResponceMessages
	Timings
}

type Ports struct {
	APIListenPort string
}
type PathVars struct {
	GetByIdParamName string
}
type ResponceMessages struct {
	ItemNotFoundMessage string
}
type Timings struct {
	APIRequestTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		Ports: Ports{
			APIListenPort: config.APIListenPort,
		},
		PathVars: PathVars{
			GetByIdParamName: config.GetByIdParamName,
		},
		ResponceMessages: ResponceMessages{
			ItemNotFoundMessage: config.ItemNotFoundMessage,
		},
		Timings: Timings{
			APIRequestTimeout: config.APIRequestTimeout,
		},
	}
}
