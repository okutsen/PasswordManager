package log

import (
	"context"
	"fmt"
)

// TODO: add metrics

type correlationIDType int
const (
	// TODO: generate id in handler, use uuid.NewRandom()
	requestID correlationIDType = iota
	clientID
)

// type LoggerWithFields interface {
// 	Logger
// 	WithField(key string, value interface{}) LoggerWithFields
// }

// Baselogger
type Logger interface{
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

func NewWithContext(logger Logger, ctx context.Context) *LoggerWithContext {
	return &LoggerWithContext{
		logger: logger,
		context: ctx,
	}
}

type LoggerWithContext struct {
	logger Logger
	context context.Context
}

func (lwc *LoggerWithContext) Info(v ...interface{}) {
	if lwc.context != nil {
		// TODO: use logrus WithFields or manually construct message
		var logMessage string
		if ctxRequestID, ok := lwc.context.Value(requestID).(string); ok {
			logMessage = fmt.Sprintf("%v		requestID=%s", v, ctxRequestID)
			// lwc.logger = lwc.logger.WithField("rqId", ctxRqId)
		}
		if ctxClientID, ok := lwc.context.Value(clientID).(string); ok {
			logMessage = fmt.Sprintf("%v		clientID=%s", v, ctxClientID)
			// lwc.logger = lwc.logger.WithField("sessionId", ctxSessionId)
		}
		lwc.logger.Info(logMessage)
	}
}