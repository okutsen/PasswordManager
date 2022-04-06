package api

import (
	"time"

	"github.com/okutsen/PasswordManager/config"
)

type Config struct {
	Ports
	PathVars
	ResponseMessages
	Timings
}

type Ports struct {
	APIListenPort string
}
type PathVars struct {
	GetByIdParamName string
}
type ResponseMessages struct {
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
		ResponseMessages: ResponseMessages{
			ItemNotFoundMessage: config.ItemNotFoundMessage,
		},
		Timings: Timings{
			APIRequestTimeout: config.APIRequestTimeout,
		},
	}
}
