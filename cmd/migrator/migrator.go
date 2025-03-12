package main

import (
	"flag"

	"go-micro-service-template/internal/app/migrator"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config/migrator.toml", "Path to configuration file")

	flag.Parse()

	if err := migrator.Run(configFile); err != nil {
		panic(err)
	}
}
