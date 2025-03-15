package manager

import (
	"time"
)

type options struct {
	host    string
	port    int
	timeout time.Duration
}

type Option func(*options)

func WithHost(in string) Option {
	return func(o *options) {
		o.host = in
	}
}

func WithPort(in int) Option {
	return func(o *options) {
		o.port = in
	}
}

func WithTimeout(in time.Duration) Option {
	return func(o *options) {
		o.timeout = in
	}
}
