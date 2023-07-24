package log

import (
	"github.com/fluxstack/fluxworks/contrib/kratos/conf"
	"github.com/fluxstack/fluxworks/log"
	"github.com/fluxstack/fluxworks/log/zaplog"
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"strings"
)

func NewLogger(zlog *zap.Logger) *log.Logger {
	return zaplog.New(zlog)
}

func NewKratosLogger(zlog *zap.Logger) klog.Logger {
	return klog.With(kzap.NewLogger(zlog))
}
func NewZapLogger(c *conf.Logger) *zap.Logger {
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

func stringToLevel(lv string) log.Level {
	switch strings.ToLower(strings.TrimSpace(lv)) {
	case "debug":
		return log.LevelDebug
	case "info":
		return log.LevelInfo
	case "warn":
		return log.LevelWarn
	case "error":
		return log.LevelError
	case "fatal":
		return log.LevelFatal
	}

	return log.LevelInfo
}
