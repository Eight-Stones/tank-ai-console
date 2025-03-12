package grpc

import (
	"strconv"

	gsrv "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"

	"go-micro-service-template/common"
	commonmocks "go-micro-service-template/common/common_mocks"
	"go-micro-service-template/internal/controller/grpc/handler/probe"
	er "go-micro-service-template/pkg/error"
)

type Option func(*options)

// defaultOptions define default option.
var defaultOptions = options{
	host:     defaultHost,
	port:     defaultPort,
	handlers: []Registrar{new(probe.Probe).Register},
	log:      &commonmocks.MockLogger{},
}

type Registrar func(server.Server) error

type options struct {
	name     string
	host     string
	port     int
	handlers []Registrar
	log      common.LoggerI
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithHandler(handler Registrar) Option {
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

func New(opt ...Option) (*Controller, error) {
	cfg := defaultOptions
	for _, o := range opt {
		o(&cfg)
	}

	service := micro.NewService(
		micro.Name(cfg.name),
		micro.Server(
			gsrv.NewServer(),
		),
		micro.Address(cfg.host+":"+strconv.Itoa(cfg.port)),
	)

	service.Init()

	for _, r := range cfg.handlers {
		if err := r(service.Server()); err != nil {
			return nil, er.Wrap(err, "error during register handlers")
		}
	}

	return &Controller{
		service: service,
		opts:    cfg,
	}, nil
}
