package example

import (
	"context"

	"go-micro-service-template/entity"

	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (e *Example) Read(ctx context.Context) ([]entity.Example, error) {
	var (
		itms []entity.Example
		err  error
	)
	err = e.opts.txp.Transact(ctx, func(tx txp.ConnectionI) error {
		itms, err = e.read(ctx, tx)
		if err != nil {
			return er.Wrap(err, "examples read failed")
		}

		return nil
	})

	if err != nil {
		return nil, er.Wrap(err, "transact")
	}

	return itms, nil
}

func (e *Example) read(ctx context.Context, tx txp.MethodI) ([]entity.Example, error) {
	exps, err := e.opts.gateway.ex.ReadExamples(ctx, tx, nil)
	if err != nil {
		return nil, er.Wrap(err, "ReadExamples")
	}

	return exps, nil
}
