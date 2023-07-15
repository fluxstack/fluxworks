package zaplog

import (
	log2 "github.com/fluxstack/fluxworks/log"
	"go.uber.org/zap"
)

func NewAdapter(log *zap.Logger) log2.Adapter {
	return &ZapLog{log: log}
}

type ZapLog struct {
	log *zap.Logger
}

func (z *ZapLog) Log(level log2.Level, fields log2.Fields) error {
	var msg string
	_msg, ok := fields[log2.DefaultMessageKey]
	if ok {
		msg, _ = _msg.(string)
	}

	var data []zap.Field = make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		if k == log2.DefaultMessageKey {
			continue
		}
		data = append(data, zap.Any(k, v))
	}

	switch level {
	case log2.LevelDebug:
		z.log.Debug(msg, data...)
	case log2.LevelInfo:
		z.log.Info(msg, data...)
	case log2.LevelWarn:
		z.log.Warn(msg, data...)
	case log2.LevelError:
		z.log.Error(msg, data...)
	case log2.LevelFatal:
		z.log.Fatal(msg, data...)
	}
	return nil
}
