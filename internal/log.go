package internal

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
	SetOutput(w io.Writer)
}

// type BaseLogger struct {
// 	 TODO: save created loggers
// 	 loggers []Logger
// }
// method of BaseLogger
func NewLogger() Logger {
	return logrus.New()
}
