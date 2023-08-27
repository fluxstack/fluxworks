package logging

import (
	"context"
	"github.com/weflux/fluxworks/types"
)

type F types.M
type Logger struct {
	adapter Adapter
}

func New(adapter Adapter) *Logger {
	return &Logger{adapter: With(adapter, types.M{})}
}

const DefaultMessageKey = "msg"
const DefaultErrorKey = "error"

func (l *Logger) With(fields types.M) *Logger {
	return &Logger{adapter: With(l.adapter, fields)}
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	return &Logger{adapter: WithContext(ctx, l.adapter)}
}

func (l *Logger) Info(msg string, fields ...types.M) {
	kvs := merge(fields...)
	kvs[DefaultMessageKey] = msg
	l.adapter.Log(LevelInfo, kvs)
}

func (l *Logger) Debug(msg string, fields ...types.M) {
	kvs := merge(fields...)
	kvs[DefaultMessageKey] = msg
	l.adapter.Log(LevelDebug, kvs)
}

func (l *Logger) Warn(msg string, fields ...types.M) {

	kvs := merge(fields...)
	kvs[DefaultMessageKey] = msg
	l.adapter.Log(LevelWarn, kvs)
}

func (l *Logger) Error(msg string, err error, fields ...types.M) {
	kvs := merge(fields...)
	kvs[DefaultMessageKey] = msg
	kvs[DefaultErrorKey] = err.Error()
	l.adapter.Log(LevelError, kvs)
}

var defaultLogger *Logger

func SetLogger(adapter Adapter) {
	defaultLogger = New(adapter)
}

func DefaultLogger() *Logger {
	return defaultLogger
}

func Debug(ctx context.Context, msg string, fields ...types.M) {
	DefaultLogger().WithContext(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...types.M) {
	DefaultLogger().WithContext(ctx).Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...types.M) {
	DefaultLogger().WithContext(ctx).Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, err error, fields ...types.M) {
	DefaultLogger().WithContext(ctx).Error(msg, err, fields...)
}
