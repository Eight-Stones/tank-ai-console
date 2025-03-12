package example

import (
	"context"

	"go-micro-service-template/entity"
	"go-micro-service-template/internal/gateway/storage/postgres/model"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (c *Client) CreateExample(ctx context.Context, tx txp.MethodI, in *entity.Example) error {
	query, args, err := buildExamplesInsert(model.ConvertExampleToModel(in))
	if err != nil {
		return er.InvalidArgumentType.Wrap(err, "build query")
	}

	row := tx.Row(ctx, query, args...)
	if err = row.Scan(&in.ID); err != nil {
		return er.InternalType.Wrap(err, "scan row")
	}

	return nil
}
