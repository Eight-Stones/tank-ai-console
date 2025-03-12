package transactioner

import (
	"time"

	ml "go-micro.dev/v4/logger"
)

type LoggerI interface {
	Fields(fields map[string]interface{}) ml.Logger
	Log(level ml.Level, v ...interface{})
	Logf(level ml.Level, format string, v ...interface{})
}

type options struct {
	host            string
	port            int
	username        string
	password        string
	dbname          string
	sslMode         string
	connMaxLifetime time.Duration
	maxOpenConns    int32
	connMaxIdleTime time.Duration
	maxIdleConns    int32

	logger LoggerI
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

func WithUsername(in string) Option {
	return func(o *options) {
		o.username = in
	}
}

func WithPassword(in string) Option {
	return func(o *options) {
		o.password = in
	}
}

func WithDBName(in string) Option {
	return func(o *options) {
		o.dbname = in
	}
}

func WithSSLMode(in string) Option {
	return func(o *options) {
		o.sslMode = in
	}
}

func WithConnMaxLifetime(in time.Duration) Option {
	return func(o *options) {
		o.connMaxLifetime = in
	}
}

func WithConnMaxIdleTime(in time.Duration) Option {
	return func(o *options) {
		o.connMaxIdleTime = in
	}
}

func WithMaxOpenConns(in int32) Option {
	return func(o *options) {
		o.maxOpenConns = in
	}
}

func WithMaxIdleConns(in int32) Option {
	return func(o *options) {
		o.maxIdleConns = in
	}
}

func WithLogger(in LoggerI) Option {
	return func(o *options) {
		o.logger = in
	}
}
