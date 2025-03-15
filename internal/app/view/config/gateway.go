package config

import (
	"go-micro.dev/v4/config"
)

type Gateway struct {
	TankClient Client
}

type Client struct {
	Host string
	Port int
}

func newGateway(cfg config.Config) Gateway {
	return Gateway{
		TankClient: newRestClientTankExample(cfg),
	}
}

func newRestClientTankExample(cfg config.Config) Client {
	return Client{
		Host: cfg.Get("gateway", "rest", "tank", "host").String(defaultRestHost),
		Port: cfg.Get("gateway", "rest", "tank", "port").Int(defaultRestPort),
	}
}
