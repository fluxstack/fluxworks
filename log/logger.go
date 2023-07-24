package log

import (
	"fmt"
)

type Logger struct {
	adapter Adapter
}

func New(adapter Adapter) *Logger {
	return &Logger{adapter: With(adapter, Fields{})}
	//return &Logger{adapter: adapter}
}

func NewWithFields(adapter Adapter, fields Fields) *Logger {
	return &Logger{adapter: With(adapter, fields)}
}

const DefaultMessageKey = "msg"

func (l *Logger) With(fields Fields) *Logger {
	return &Logger{adapter: With(l.adapter, fields)}
}

func (l *Logger) Info(args ...interface{}) {
	l.adapter.Log(LevelInfo, Fields{
		DefaultMessageKey: fmt.Sprint(args...),
	})
}

func (l *Logger) Infow(fields Fields) {
	l.adapter.Log(LevelInfo, fields)
}

func (l *Logger) Debug(args ...interface{}) {
	l.adapter.Log(LevelDebug, Fields{
		DefaultMessageKey: fmt.Sprint(args...),
	})
}

func (l *Logger) Debugw(fields Fields) {
	l.adapter.Log(LevelDebug, fields)
}

func (l *Logger) Warn(args ...interface{}) {
	l.adapter.Log(LevelWarn, Fields{
		DefaultMessageKey: fmt.Sprint(args...),
	})
}

func (l *Logger) Warnw(fields Fields) {
	l.adapter.Log(LevelWarn, fields)
}

func (l *Logger) Error(args ...interface{}) {
	l.adapter.Log(LevelError, Fields{
		DefaultMessageKey: fmt.Sprint(args...),
	})
}

func (l *Logger) Errorw(fields Fields) {
	l.adapter.Log(LevelError, fields)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.adapter.Log(LevelFatal, Fields{
		DefaultMessageKey: fmt.Sprint(args...),
	})
}

func (l *Logger) Fatalw(fields Fields) {
	l.adapter.Log(LevelFatal, fields)
}
