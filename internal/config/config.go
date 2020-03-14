package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initialize logger configuration for micro service
func BuildLogger() (*zap.SugaredLogger, error) {
	configLogger := zap.NewDevelopmentConfig()
	configLogger.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := configLogger.Build(zap.AddCaller())
	if err != nil {
		return nil, err
	}
	logger := l.Sugar()
	defer l.Sync()
	return logger, nil
}
