package log

import (
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/weflux/fluxworks/logging"
	"github.com/weflux/fluxworks/logging/zaplog"
	"go.uber.org/zap"
	"strings"
)

func NewLogger(zlog *zap.Logger) *logging.Logger {
	return zaplog.New(zlog)
}

func NewKratosLogger(zlog *zap.Logger) klog.Logger {
	return klog.With(kzap.NewLogger(zlog))
}

type Config struct {
	Logger string    `json:"logger"`
	Level  string    `json:"level"`
	Zap    ZapConfig `json:"zap"`
}

type ZapConfig struct {
	Production bool `json:"production"`
}

func NewZapLogger(c Config) *zap.Logger {
	zlog, err := zaplog.NewZapLogger(zaplog.Options{
		CallerSkip: 3,
		Production: c.Zap.Production,
		Level:      stringToLevel(c.Level),
	})
	if err != nil {
		panic(err)
	}
	return zlog
}

func stringToLevel(lv string) logging.Level {
	switch strings.ToLower(strings.TrimSpace(lv)) {
	case "debug":
		return logging.LevelDebug
	case "info":
		return logging.LevelInfo
	case "warn":
		return logging.LevelWarn
	case "error":
		return logging.LevelError
	case "fatal":
		return logging.LevelFatal
	}

	return logging.LevelInfo
}
