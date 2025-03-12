package transactioner

import (
	"context"
)

type Row interface {
	Scan(dest ...any) error
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Txer interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type MethodI interface {
	Exec(ctx context.Context, sql string, args ...any) error
	Rows(ctx context.Context, sql string, args ...any) (Rows, func(), error)
	Row(ctx context.Context, sql string, args ...any) Row
}

type ConnectionI interface {
	MethodI
	Txer
}
