package config

import (
	"go-micro.dev/v4/config"
)

type App struct {
	Name string
}

func newApp(cfg config.Config) App {
	return App{
		Name: cfg.Get("app", "name").String("example"),
	}
}
