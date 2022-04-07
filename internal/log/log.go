package log

// TODO: add metrics

type Logger interface{
	BaseLogger
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

type BaseLogger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

// type LoggerWithContext struct {
// 	logger Logger
// 	context context.Context
// }

// func (lwc *LoggerWithContext) 