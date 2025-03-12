package example

import (
	"context"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/gateway/storage/postgres/model"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (c *Client) UpdateExample(ctx context.Context, tx txp.MethodI, in *entity.Example) error {
	query, args, err := buildExamplesUpdate(model.ConvertExampleToModel(in))
	if err != nil {
		return er.InvalidArgumentType.Wrap(err, "build query")
	}

	if err = tx.Exec(ctx, query, args...); err != nil {
		return er.InternalType.Wrap(err, "Exec")
	}

	return nil
}
