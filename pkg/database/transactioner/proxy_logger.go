package transactioner

import (
	"context"
	"strings"

	"go-micro.dev/v4/logger"
)

const (
	queryKey = "query"
	argsKey  = "args"
)

func clearQuery(query string) string {
	query = strings.ReplaceAll(query, "\n", "")
	query = strings.ReplaceAll(query, "\t", "")
	return query
}

type ProxyLogger struct {
	logger LoggerI
	conn   ConnectionI
}

func newProxyLogger(logger LoggerI, conn ConnectionI) *ProxyLogger {
	return &ProxyLogger{
		logger: logger,
		conn:   conn,
	}
}

func (c *ProxyLogger) Commit(ctx context.Context) error {
	if err := c.conn.Commit(ctx); err != nil {
		c.logger.Logf(logger.DebugLevel, "commit error: %v", err)
		return err
	}
	return nil
}

func (c *ProxyLogger) Rollback(ctx context.Context) error {
	if err := c.conn.Rollback(ctx); err != nil {
		c.logger.Logf(logger.DebugLevel, "rollback error: %v", err)
		return err
	}
	return nil
}

func (c *ProxyLogger) Exec(ctx context.Context, query string, args ...any) error {
	if err := c.conn.Exec(ctx, query, args...); err != nil {
		c.logger.Logf(logger.DebugLevel, "exec error: %v", err)
		return err
	}
	c.logger.
		Fields(map[string]any{
			queryKey: clearQuery(query),
			argsKey:  args,
		}).
		Logf(logger.DebugLevel, "exec success")
	return nil
}

func (c *ProxyLogger) Row(ctx context.Context, query string, args ...any) Row {
	c.logger.
		Fields(map[string]any{
			queryKey: clearQuery(query),
			argsKey:  args,
		}).
		Logf(logger.DebugLevel, "row success")
	return c.conn.Row(ctx, query, args...)
}

func (c *ProxyLogger) Rows(ctx context.Context, query string, args ...any) (Rows, func(), error) {
	rows, closeFn, err := c.conn.Rows(ctx, query, args...)
	if err != nil {
		c.logger.Logf(logger.DebugLevel, "query error: %v", err)
		return nil, closeFn, err
	}
	c.logger.
		Fields(map[string]any{
			queryKey: clearQuery(query),
			argsKey:  args,
		}).
		Logf(logger.DebugLevel, "row query")
	return rows, closeFn, nil
}
