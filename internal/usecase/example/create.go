package example

import (
	"context"

	"go-micro-service-template/entity"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (e *Example) Create(ctx context.Context, in *entity.Example) (int64, error) {
	var (
		id  int64
		err error
	)
	err = e.opts.txp.Transact(ctx, func(tx txp.ConnectionI) error {
		id, err = e.create(ctx, tx, in)
		if err != nil {
			return er.Wrap(err, "example create failed")
		}

		return nil
	})
	if err != nil {
		return 0, er.Wrap(err, "transact")
	}

	return id, nil
}

func (e *Example) create(ctx context.Context, tx txp.MethodI, in *entity.Example) (int64, error) {
	if err := e.opts.gateway.ex.CreateExample(ctx, tx, in); err != nil {
		return 0, er.Wrap(err, "e.opts.gateway.ex.CreateExample")
	}

	return in.ID, nil
}
