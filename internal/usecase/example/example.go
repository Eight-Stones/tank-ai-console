package example

import (
	"context"

	"go-micro-service-template/entity"

	"go-micro-service-template/pkg/database/option"
	txp "go-micro-service-template/pkg/database/transactioner"
)

type Exampler interface {
	GetExample(ctx context.Context, tx txp.MethodI, opts *option.Option) (*entity.Example, error)
	ReadExamples(ctx context.Context, tx txp.MethodI, opts *option.Option) ([]entity.Example, error)
	CreateExample(ctx context.Context, tx txp.MethodI, in *entity.Example) error
	UpdateExample(ctx context.Context, tx txp.MethodI, in *entity.Example) error
	DeleteExample(ctx context.Context, tx txp.MethodI, opts *option.Option) error
}

type Gateways struct {
	ex Exampler
}

type Example struct {
	opts options
}

func New(opts ...Option) *Example {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Example{
		opts: cfg,
	}
}
