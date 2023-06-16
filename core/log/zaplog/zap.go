package zaplog

import (
	"github.com/fluxstack/fluxworks/core/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Level       log.Level
	Output      string
	ErrorOutput string
	Production  bool
}

func New(opts Options) (*zap.Logger, error) {

	var config zap.Config
	if opts.Production {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	options := []zap.Option{
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	}

	if !opts.Production {
		options = append(options, zap.Development())
	}

	level := zapcore.DebugLevel
	switch opts.Level {
	case log.LevelFatal:
		level = zapcore.FatalLevel
	case log.LevelError:
		level = zapcore.ErrorLevel
	case log.LevelWarn:
		level = zapcore.WarnLevel
	case log.LevelInfo:
		level = zapcore.InfoLevel
	}

	config.Level = zap.NewAtomicLevelAt(level)

	if opts.Output != "" {
		config.OutputPaths = append(config.OutputPaths, opts.Output)
	}

	if opts.ErrorOutput != "" {
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, opts.ErrorOutput)
	}

	return config.Build(options...)
}

func ProductionLogger() *log.Logger {
	l, err := New(Options{Production: true})
	if err != nil {
		panic(err)
	}
	return log.New(NewAdapter(l))
}

func DebugLogger() *log.Logger {
	l, err := New(Options{Production: false})
	if err != nil {
		panic(err)
	}
	return log.New(NewAdapter(l))
}
