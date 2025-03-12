package example

import (
	"context"

	"go-micro-service-template/entity"
	model2 "go-micro-service-template/internal/gateway/storage/postgres/model"
	"go-micro-service-template/pkg/database/option"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (c *Client) ReadExamples(ctx context.Context, tx txp.MethodI, opts *option.Option) ([]entity.Example, error) {
	exps, err := examples(ctx, tx, opts)
	if err != nil {
		return nil, er.Wrap(err, "read examples")
	}

	return model2.ConvertModelsToExamples(exps), nil
}

func examples(ctx context.Context, tx txp.MethodI, opts *option.Option) ([]model2.Example, error) {
	query, args, err := buildExamplesQuery(opts)
	if err != nil {
		return nil, er.InvalidArgumentType.Wrap(err, "build query")
	}

	rows, closeFn, err := tx.Rows(ctx, query, args...)
	if err != nil {
		return nil, er.InternalType.Wrap(err, "query rows")
	}
	defer closeFn()

	var exps []model2.Example
	for rows.Next() {
		var exp model2.Example
		err = rows.Scan(
			&exp.ID,
			&exp.Code,
			&exp.Name,
			&exp.Meta,
		)
		if err != nil {
			return nil, er.InternalType.Wrap(err, "scan row")
		}
		exps = append(exps, exp)
	}

	return exps, nil
}
