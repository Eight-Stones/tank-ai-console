package example

import (
	"context"

	"go-micro-service-template/entity"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (e *Example) Update(ctx context.Context, in *entity.Example) error {
	var (
		err error
	)
	err = e.opts.txp.Transact(ctx, func(tx txp.ConnectionI) error {
		err = e.update(ctx, tx, in)
		if err != nil {
			return er.Wrap(err, "example update failed")
		}

		return nil
	})
	if err != nil {
		return er.Wrap(err, "transact")
	}

	return nil
}

func (e *Example) update(ctx context.Context, tx txp.MethodI, in *entity.Example) error {
	if err := e.opts.gateway.ex.UpdateExample(ctx, tx, in); err != nil {
		return er.Wrap(err, "e.opts.gateway.ex.UpdateExample")
	}

	return nil
}
