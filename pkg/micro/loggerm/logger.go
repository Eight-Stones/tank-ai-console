package loggerm

import (
	"fmt"

	mzap "github.com/go-micro/plugins/v4/logger/zap"

	"go-micro.dev/v4/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(opts ...Option) (logger.Logger, error) {
	cfg := defaultOptions
	for _, opt := range opts {
		opt(&cfg)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Development:       cfg.Development,
		DisableCaller:     cfg.DisableCaller,
		DisableStacktrace: cfg.DisableStacktrace,
		Encoding:          cfg.Encoding,
		EncoderConfig:     encoderCfg,
		OutputPaths:       cfg.OutputPaths,
		ErrorOutputPaths:  cfg.ErrorOutputPaths,
		InitialFields:     cfg.InitialFields,
	}

	l, err := mzap.NewLogger(
		mzap.WithConfig(config),
		logger.WithCallerSkipCount(0),
		logger.WithLevel(logger.Level(cfg.Level)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to created new micro zap: %w", err)
	}

	logger.DefaultLogger = l

	return l, nil
}
