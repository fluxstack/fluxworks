package zaplog

import (
	log2 "github.com/fluxstack/fluxworks/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Level       log2.Level
	Output      string
	ErrorOutput string
	Production  bool
	CallerSkip  int
}

func NewZapLogger(opt Options) (*zap.Logger, error) {

	var config zap.Config
	if opt.Production {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	if opt.CallerSkip == 0 {
		opt.CallerSkip = 3
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	options := []zap.Option{
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(opt.CallerSkip),
	}

	if !opt.Production {
		options = append(options, zap.Development())
	}

	level := zapcore.DebugLevel
	switch opt.Level {
	case log2.LevelFatal:
		level = zapcore.FatalLevel
	case log2.LevelError:
		level = zapcore.ErrorLevel
	case log2.LevelWarn:
		level = zapcore.WarnLevel
	case log2.LevelInfo:
		level = zapcore.InfoLevel
	}

	config.Level = zap.NewAtomicLevelAt(level)

	if opt.Output != "" {
		config.OutputPaths = append(config.OutputPaths, opt.Output)
	}

	if opt.ErrorOutput != "" {
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, opt.ErrorOutput)
	}

	return config.Build(options...)
}

func New(zlog *zap.Logger) *log2.Logger {
	return log2.New(NewAdapter(zlog))
}
