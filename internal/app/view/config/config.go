package config

import (
	er "go-micro-service-template/pkg/error"
	"go-micro-service-template/pkg/micro/configm"
)

type Config struct {
	App           App
	Observability Observability
	Controller    Controller
	Gateway       Gateway
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
		Gateway:       newGateway(cfg),
	}, nil
}
