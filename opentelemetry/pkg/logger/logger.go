package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", "./logs/fresh-market.log"},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, _ = config.Build()
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
	err := log.Sync()
	if err != nil {
		return
	}
}

func Error(msg string, errors error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", errors))

	log.Error(msg, fields...)
	err := log.Sync()
	if err != nil {
		return
	}
}
