package infra

import (
	"github.com/ian0113/go-gin-mvc/config"

	"go.uber.org/zap"
)

var (
	globalLogger *zap.Logger
)

func NewLogger(cfg *config.Config) *zap.Logger {
	var newLoggerFunc func(options ...zap.Option) (*zap.Logger, error)
	switch cfg.App.Mode {
	case config.AppModeDevelopment:
		newLoggerFunc = zap.NewDevelopment
	case config.AppModeProduction:
		newLoggerFunc = zap.NewProduction
	default:
		newLoggerFunc = zap.NewProduction
	}
	logger, err := newLoggerFunc()
	if err != nil {
		panic(err)
	}
	return logger.Named(cfg.App.Name)
}

func InitLogger(cfg *config.Config) {
	globalLogger = NewLogger(cfg)
}

func GetLogger() *zap.Logger {
	return globalLogger
}
