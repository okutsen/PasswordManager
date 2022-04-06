package main

import "github.com/okutsen/PasswordManager/internal/server"

func main() {
	domain := server.NewServer()
	domain.Start()
}
