package config

import (
	er "go-micro-service-template/pkg/error"
	"go-micro-service-template/pkg/micro/configm"
)

type Config struct {
	Storage Storage
}

func New(path string) (*Config, error) {
	cfg, err := configm.New(path)
	if err != nil {
		return nil, er.Wrap(err, "new config")
	}

	return &Config{
		Storage: NewStorage(cfg),
	}, nil
}
