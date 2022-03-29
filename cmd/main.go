package main

import "github.com/okutsen/PasswordManager/internal"

// TODO: write Logger interface on Logrus Logger
// type Logger interface {
// 	Info()
// }
// log Logger := &log.Logger

// TODO: password tips or reset questions
func main() {
	clientServer := &internal.ClientAPI{}
	clientServer.Start()
}
