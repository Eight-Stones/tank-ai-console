package example

import (
	"context"

	"go-micro-service-template/pkg/database/option"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (e *Example) Delete(ctx context.Context, id int64) error {
	var (
		err error
	)
	err = e.opts.txp.Transact(ctx, func(tx txp.ConnectionI) error {
		err = e.delete(ctx, tx, id)
		if err != nil {
			return er.Wrap(err, "example delete failed")
		}

		return nil
	})
	if err != nil {
		return er.Wrap(err, "transact")
	}

	return nil
}

func (e *Example) delete(ctx context.Context, tx txp.MethodI, id int64) error {
	if err := e.opts.gateway.ex.DeleteExample(ctx, tx, &option.Option{
		Filter: option.NewExpression("id", option.EQ, id),
	}); err != nil {
		return er.Wrap(err, "e.opts.gateway.ex.DeleteExample")
	}

	return nil
}
