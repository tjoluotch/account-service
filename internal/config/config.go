package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Logger struct {
	*zap.SugaredLogger
}

// initialize server logger level encoder and configuration
func (l *Logger) Set() *Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
		return nil
	}
	logger.Info("logger initialised")
	defer logger.Sync()
	loggerType := Logger{logger.Sugar()}
	return &loggerType
}
