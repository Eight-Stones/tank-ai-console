package config

import (
	"go-micro.dev/v4/config"
)

type Observability struct {
	Logger Logger
}

type Logger struct {
	Level string
	Keys  map[string]any
}

func newObservability(cfg config.Config) Observability {
	return Observability{
		Logger: NewLogger(cfg),
	}
}

func NewLogger(cfg config.Config) Logger {
	keys := make(map[string]any)
	_ = cfg.Get("logger", "keys").Scan(&keys)

	return Logger{
		Level: cfg.Get("logger", "level").String("debug"),
		Keys:  keys,
	}
}
