package example

import (
	"context"
	"database/sql"
	"errors"

	"go-micro-service-template/entity"
	model2 "go-micro-service-template/internal/gateway/storage/postgres/model"
	"go-micro-service-template/pkg/database/option"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (c *Client) GetExample(ctx context.Context, tx txp.MethodI, opts *option.Option) (*entity.Example, error) {
	exp, err := example(ctx, tx, opts)
	if err != nil {
		return nil, er.Wrap(err, "get example")
	}

	return model2.ConvertModelToExample(exp), nil
}

func example(ctx context.Context, tx txp.MethodI, opts *option.Option) (*model2.Example, error) {
	query, args, err := buildExamplesQuery(opts)
	if err != nil {
		return nil, er.InvalidArgumentType.Wrap(err, "build query")
	}

	var exp model2.Example
	row := tx.Row(ctx, query, args...)
	err = row.Scan(
		&exp.ID,
		&exp.Code,
		&exp.Name,
		&exp.Meta,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, er.NotFoundType.Wrap(err, "example not found")
		}
		return nil, er.InternalType.Wrap(err, "scan row")
	}

	return &exp, nil
}
