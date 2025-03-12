package example

import (
	"context"

	"go-micro-service-template/pkg/database/option"
	txp "go-micro-service-template/pkg/database/transactioner"
	er "go-micro-service-template/pkg/error"
)

func (c *Client) DeleteExample(ctx context.Context, tx txp.MethodI, opts *option.Option) error {
	query, args, err := buildExamplesDelete(opts)
	if err != nil {
		return er.InvalidArgumentType.Wrap(err, "build query")
	}

	if err = tx.Exec(ctx, query, args...); err != nil {
		return er.InternalType.Wrap(err, "exec")
	}

	return nil
}
