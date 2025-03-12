package transactioner

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Tx struct {
	tx pgx.Tx
}

func (t *Tx) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *Tx) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

func (t *Tx) Exec(ctx context.Context, query string, args ...any) error {
	res, err := t.tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (t *Tx) Row(ctx context.Context, query string, args ...any) Row {
	return t.tx.QueryRow(ctx, query, args...)
}

func (t *Tx) Rows(ctx context.Context, query string, args ...any) (Rows, func(), error) {
	rows, err := t.tx.Query(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}

	fn := func() { rows.Close() }

	return rows, fn, nil
}
