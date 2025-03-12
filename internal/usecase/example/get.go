package example

import (
	"context"

	"go-micro-service-template/entity"
	"go-micro-service-template/pkg/database/option"

	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (e *Example) Get(ctx context.Context, id int64) (*entity.Example, error) {
	var (
		exp *entity.Example
		err error
	)
	err = e.opts.txp.Transact(ctx, func(tx txp.ConnectionI) error {
		exp, err = e.get(ctx, tx, id)
		if err != nil {
			return er.Wrap(err, "example get failed")
		}

		return nil
	})
	if err != nil {
		return nil, er.Wrap(err, "transact")
	}

	return exp, nil
}

func (e *Example) get(ctx context.Context, tx txp.MethodI, id int64) (*entity.Example, error) {
	exp, err := e.opts.gateway.ex.GetExample(ctx, tx, &option.Option{
		Filter: option.NewExpression("id", option.EQ, id),
		Limit:  1,
	})
	if err != nil {
		return nil, er.Wrap(err, "GetExample")
	}

	return exp, nil
}
