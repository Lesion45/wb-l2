package logger

import "go.uber.org/zap"

func NewAppZap() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	return zap.Must(cfg.Build())
}

func NewReqZap() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()

	return zap.Must(cfg.Build())
}
