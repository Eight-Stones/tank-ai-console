package config

import (
	"go-micro.dev/v4/config"
)

const (
	defaultRestHost = "localhost"
	defaultRestPort = 8080
	defaultGrpcHost = "localhost"
	defaultGrpcPort = 50051
)

type Controller struct {
	ExampleRest RestBook
	ExampleGrpc RestBook
}

type RestBook struct {
	Host string
	Port int
}

func newController(cfg config.Config) Controller {
	return Controller{
		ExampleRest: newRestExample(cfg),
		ExampleGrpc: newGrpcExample(cfg),
	}
}

func newRestExample(cfg config.Config) RestBook {
	return RestBook{
		Host: cfg.Get("controller", "rest", "host").String(defaultRestHost),
		Port: cfg.Get("controller", "rest", "port").Int(defaultRestPort),
	}
}

func newGrpcExample(cfg config.Config) RestBook {
	return RestBook{
		Host: cfg.Get("controller", "grpc", "host").String(defaultGrpcHost),
		Port: cfg.Get("controller", "grpc", "port").Int(defaultGrpcPort),
	}
}
