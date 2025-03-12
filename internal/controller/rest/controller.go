package rest

import (
	"strconv"

	"go-micro.dev/v4"
	"go-micro.dev/v4/codec/json"
	"go-micro.dev/v4/server"

	"go-micro-service-template/common"
	commonmocks "go-micro-service-template/common/common_mocks"
	"go-micro-service-template/internal/controller/rest/handler/probe"
	"go-micro-service-template/internal/controller/rest/middleware"
	"go-micro-service-template/pkg/micro/fastm"
)

type Option func(*options)

// defaultOptions define default option.
var defaultOptions = options{
	host:     defaultHost,
	port:     defaultPort,
	handlers: []any{new(probe.Probes)},
	log:      &commonmocks.MockLogger{},
}

type options struct {
	name     string
	host     string
	port     int
	handlers []any
	log      common.LoggerI
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithHandler(handler any) Option {
	return func(o *options) {
		o.handlers = append(o.handlers, handler)
	}
}

func WithHost(host string) Option {
	return func(o *options) {
		o.host = host
	}
}

func WithPort(port int) Option {
	return func(o *options) {
		o.port = port
	}
}

func WithLogger(logger common.LoggerI) Option {
	return func(o *options) {
		o.log = logger
	}
}

type Controller struct {
	service micro.Service
	opts    options
}

func New(opt ...Option) *Controller {
	cfg := defaultOptions
	for _, o := range opt {
		o(&cfg)
	}

	opts := []micro.Option{
		micro.Name("example"),
		micro.Server(
			fastm.NewServer(
				server.Codec("application/json", json.NewCodec),
				server.WrapHandler(middleware.ResolverWrapper),
				server.WrapHandler(middleware.PanicWrapper),
				server.WrapHandler(fastm.LogWrapper),
				server.WrapHandler(fastm.RequestIDAcquirer),
			),
		),
		micro.Address(cfg.host + ":" + strconv.Itoa(cfg.port)),
	}

	for _, h := range cfg.handlers {
		opts = append(opts, micro.Handle(h))
	}

	// Create a new service
	service := micro.NewService(
		opts...,
	)

	return &Controller{
		service: service,
		opts:    cfg,
	}
}
