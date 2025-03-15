package main

import (
	"flag"
	"fmt"
	"os"

	"go-micro-service-template/internal/app/view"
)

var (
	release   string // nolint
	buildDate string // nolint
	gitHash   string // nolint
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config/config.toml", "Path to configuration file")

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(release, buildDate, gitHash)
		return
	}

	if err := view.Run(configFile); err != nil {
		panic(err)
	}
}
