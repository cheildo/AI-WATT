package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Init initializes the global zap logger. Call once at startup.
func Init(env string) {
	var cfg zap.Config
	if env == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var err error
	log, err = cfg.Build()
	if err != nil {
		panic("logger.Init: failed to build logger: " + err.Error())
	}
}

// L returns the global logger. Panics if Init was not called.
func L() *zap.Logger {
	if log == nil {
		panic("logger.L: logger not initialized — call logger.Init first")
	}
	return log
}

// Sync flushes buffered log entries. Call on graceful shutdown.
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
