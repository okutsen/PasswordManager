package main

import "github.com/okutsen/PasswordManager/internal/api"

// TODO: password tips or reset questions

func main() {
	api := api.NewAPI()
	api.Start()
}
