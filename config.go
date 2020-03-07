package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func setLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
		return nil
	}
	logger.Info("logger initialised")
	defer logger.Sync()
	return logger.Sugar()
}
