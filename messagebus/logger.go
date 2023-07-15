package messagebus

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/fluxstack/fluxworks/log"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(logger *log.Logger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	f := log.Fields{}
	f[log.DefaultMessageKey] = msg
	f["error"] = err.Error()
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Errorw(f)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	f := log.Fields{}
	f[log.DefaultMessageKey] = msg
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Infow(f)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	f := log.Fields{}
	f[log.DefaultMessageKey] = msg
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Debugw(f)
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	f := log.Fields{}
	f[log.DefaultMessageKey] = msg
	for k, v := range fields {
		f[k] = v
	}
	l.logger.Debugw(f)
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	f := log.Fields{}
	for k, v := range fields {
		f[k] = v
	}
	logger := l.logger.With(f)
	return &Logger{logger: logger}
}
