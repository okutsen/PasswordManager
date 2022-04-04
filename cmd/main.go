package main

import (
	"github.com/okutsen/PasswordManager/internal"
)

// TODO: password tips or reset questions

func main() {
	domain := internal.NewDomainServer()
	go domain.Start()

	clientServer := internal.NewClientAPI()
	clientServer.Start()
}
