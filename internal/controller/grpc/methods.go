package grpc

import (
	"errors"
	"net/http"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
)

var _ micro.Service = &Controller{}

func (c Controller) Name() string {
	return c.service.Name()
}

func (c Controller) Init(option ...micro.Option) {
	c.service.Init(option...)
}

func (c Controller) Options() micro.Options {
	return c.service.Options()
}

func (c Controller) Client() client.Client {
	return c.service.Client()
}

func (c Controller) Server() server.Server {
	return c.service.Server()
}

func (c Controller) Run() error {
	return c.service.Run()
}

func (c Controller) String() string {
	return c.service.Name()
}

func (c Controller) Start() {
	go func() {
		c.opts.log.Infof("grpc server '%v' started on '%v:%d'", c.opts.name, c.opts.host, c.opts.port)
		if err := c.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.opts.log.Errorf("error during run grpc: %v", err)
			return
		}
		c.opts.log.Infof("rest server '%v' stopped", c.opts.name)
	}()
}
