package config

import "time"

// TODO: use viper lib

// API
// Host
const (
	APIHostURL string = "http://localhost"
)

// Ports
const (
	APIListenPort = "10000"
)

// PathVars
const (
	GetByIdParamName = "recordName"
)

// ResponceMessages
const (
	ItemNotFoundMessage = "Item not found"
)

// Timings
const (
	APIRequestTimeout = 60 * time.Second
)

// Server
// Host
const (
	ServerHostURL = "http://localhost"
)

// Ports
const (
	ServerListenPort = "10001"
)

// Timings
const (
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
)

// TODO: more fields (dateCreated, url, description)
// type Record struct {
// 	id       int
// 	login    string
// 	password string
// }
