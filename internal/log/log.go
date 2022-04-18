package log

import (
	"bytes"
	"fmt"
)

// TODO: add metrics
type Logger interface{
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
}

type BaseLogger struct {
	
}






type Fields map[string]any

type LoggerWithContext struct {
	logger Logger
	fields Fields
	prefix string
	sep string
}

func NewWithFields(logger Logger) *LoggerWithContext {
	return &LoggerWithContext{
		logger: logger,
		prefix: " ",
		sep: "\t",
	}
}

func (lwc *LoggerWithContext) SetPrefix(prefix string) {
	lwc.prefix = prefix
}

func (lwc *LoggerWithContext) SetSep(sep string) {
	lwc.sep = sep
}

func (lwc *LoggerWithContext) WithFields(fields Fields) *LoggerWithContext {
	lwc.addFields(fields)
	return lwc
}

// to pkg?
func (lwc *LoggerWithContext) addFields(f Fields) {
	// concat maps?
	for k, v := range f {
		lwc.fields[k] = v
	}
}

func (lwc *LoggerWithContext) ComposeMessage(v []any) string {
	return lwc.composeMessage(fmt.Sprint(v...))
}

func (lwc *LoggerWithContext) ComposeMessagef(format string, v []any) string {
	return lwc.composeMessage(fmt.Sprintf(format, v...))
}

// to pkg?
func (lwc *LoggerWithContext) composeMessage(initial string) string {
	var buff bytes.Buffer
	buff.WriteString(initial)
	buff.WriteString(lwc.prefix)
    for key, value := range lwc.fields {
        fmt.Fprintf(&buff, "%s%s=%v", lwc.sep, key, value)
    }
    return buff.String()
}

// any way not to override every method? 
func (lwc *LoggerWithContext) Print(v ...any) {
	lwc.logger.Info(lwc.ComposeMessage(v))
}

func (lwc *LoggerWithContext) Printf(format string, v ...any) {
	lwc.logger.Info(lwc.ComposeMessagef(format, v))
}
