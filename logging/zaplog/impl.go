package zaplog

import (
	"github.com/weflux/fluxworks/logging"
	"github.com/weflux/fluxworks/types"
	"go.uber.org/zap"
)

func NewAdapter(log *zap.Logger) logging.Adapter {
	return &ZapLog{log: log}
}

type ZapLog struct {
	log *zap.Logger
}

func (z *ZapLog) Log(level logging.Level, fields types.M) error {
	var msg string
	_msg, ok := fields[logging.DefaultMessageKey]
	if ok {
		msg, _ = _msg.(string)
	}

	var data []zap.Field = make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		if k == logging.DefaultMessageKey {
			continue
		}
		data = append(data, zap.Any(k, v))
	}

	switch level {
	case logging.LevelDebug:
		z.log.Debug(msg, data...)
	case logging.LevelInfo:
		z.log.Info(msg, data...)
	case logging.LevelWarn:
		z.log.Warn(msg, data...)
	case logging.LevelError:
		z.log.Error(msg, data...)
	case logging.LevelFatal:
		z.log.Fatal(msg, data...)
	}
	return nil
}
