package config

// TODO: use viper lib

const (
	ClientAPIPort    string = "10000"
	DomainServerPort string = "10001"
	DomainServerURL  string = "http://localhost"
)

// TODO: more fields (dateCreated, url, description)
type Record struct {
	id       int
	login    string
	password string
}
