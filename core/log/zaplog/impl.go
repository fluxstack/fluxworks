package zaplog

import (
	"github.com/fluxstack/fluxworks/core/log"
	"go.uber.org/zap"
)

func NewAdapter(log *zap.Logger) log.Adapter {
	return &ZapLog{log: log}
}

type ZapLog struct {
	log *zap.Logger
}

func (z *ZapLog) Log(level log.Level, fields log.Fields) error {
	var msg string
	_msg, ok := fields[log.DefaultMessageKey]
	if ok {
		msg, _ = _msg.(string)
	}

	var data []zap.Field = make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		if k == log.DefaultMessageKey {
			continue
		}
		data = append(data, zap.Any(k, v))
	}

	switch level {
	case log.LevelDebug:
		z.log.Debug(msg, data...)
	case log.LevelInfo:
		z.log.Info(msg, data...)
	case log.LevelWarn:
		z.log.Warn(msg, data...)
	case log.LevelError:
		z.log.Error(msg, data...)
	case log.LevelFatal:
		z.log.Fatal(msg, data...)
	}
	return nil
}
