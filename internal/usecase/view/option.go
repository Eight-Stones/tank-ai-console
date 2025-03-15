package view

import (
	"time"
)

type options struct {
	redrawTimeout time.Duration
	manager       TankManager
}

type Option func(*options)

func WithReDrawTimeout(in time.Duration) Option {
	return func(o *options) {
		o.redrawTimeout = in
	}
}

func WithManager(in TankManager) Option {
	return func(o *options) {
		o.manager = in
	}
}
