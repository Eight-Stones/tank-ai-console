package config

import (
	er "go-micro-service-template/pkg/error"
	"go-micro-service-template/pkg/micro/configm"
)

type Config struct {
	App           App
	Observability Observability
	Controller    Controller
	Storage       Storage
}

func New(path string) (*Config, error) {
	cfg, err := configm.New(path)
	if err != nil {
		return nil, er.Wrap(err, "new config")
	}

	return &Config{
		App:           newApp(cfg),
		Observability: newObservability(cfg),
		Controller:    newController(cfg),
		Storage:       newStorage(cfg),
	}, nil
}
