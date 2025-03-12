package configm

import (
	"github.com/go-micro/plugins/v4/config/encoder/toml"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
)

// New initialize toml config by using go-micro config.
func New(path string) (config.Config, error) {
	enc := toml.NewEncoder()

	r := config.WithReader(
		json.NewReader(
			reader.WithEncoder(
				enc,
			),
		),
	)

	s := config.WithSource(
		file.NewSource(
			source.WithEncoder(enc),
			file.WithPath(path),
		),
	)

	cfg, err := config.NewConfig(
		r,
		s,
	)
	if err != nil {
		return nil, err
	}

	if err = cfg.Load(); err != nil {
		return nil, err
	}

	config.DefaultConfig = cfg

	return cfg, nil
}
