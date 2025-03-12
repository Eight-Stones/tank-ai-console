package transactioner

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn struct {
	conn *pgxpool.Conn
}

func (c *Conn) Commit(_ context.Context) error {
	return nil
}

func (c *Conn) Rollback(_ context.Context) error {
	return nil
}

func (c *Conn) Exec(ctx context.Context, query string, args ...any) error {
	_, err := c.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) Row(ctx context.Context, query string, args ...any) Row {
	return c.conn.QueryRow(ctx, query, args...)
}

func (c *Conn) Rows(ctx context.Context, query string, args ...any) (Rows, func(), error) {
	rows, err := c.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}

	fn := func() { rows.Close() }

	return rows, fn, nil
}
