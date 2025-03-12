package transactioner

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	er "go-micro-service-template/pkg/error"
)

type TxOptions struct {
	IsoLevel       string
	AccessMode     string
	DeferrableMode string
	BeginQuery     string
}

type TxProvider struct {
	pool *pgxpool.Pool
	cfg  options
}

func New(opts ...Option) *TxProvider {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	return &TxProvider{
		cfg: cfg,
	}
}

func (t *TxProvider) Connect(ctx context.Context) error {
	connString := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v sslmode=%v pool_max_conns=%v",
		t.cfg.username,
		t.cfg.password,
		t.cfg.host,
		t.cfg.port,
		t.cfg.dbname,
		t.cfg.sslMode,
		t.cfg.maxOpenConns,
	)

	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return er.Wrap(err, "ParseConfig")
	}

	cfg.MaxConns = t.cfg.maxOpenConns
	cfg.MaxConnIdleTime = t.cfg.connMaxIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return er.Wrap(err, "NewWithConfig")
	}

	if err = pool.Ping(ctx); err != nil {
		return er.Wrap(err, "Ping")
	}

	t.pool = pool

	return nil
}

func (t *TxProvider) Connection(ctx context.Context) (ConnectionI, error) {
	conn, err := t.pool.Acquire(ctx)
	if err != nil {
		return nil, er.Wrap(err, "pool acquire")
	}
	if t.cfg.logger != nil {
		return newProxyLogger(t.cfg.logger, &Conn{conn: conn}), nil
	}

	return &Conn{conn: conn}, nil
}

func (t *TxProvider) Tx(ctx context.Context, opts ...TxOptions) (ConnectionI, error) {
	opt := pgx.TxOptions{}
	if len(opts) > 0 {
		opt = pgx.TxOptions{
			IsoLevel:       pgx.TxIsoLevel(opts[0].IsoLevel),
			AccessMode:     pgx.TxAccessMode(opts[0].AccessMode),
			DeferrableMode: pgx.TxDeferrableMode(opts[0].DeferrableMode),
			BeginQuery:     opts[0].BeginQuery,
		}
	}

	tx, err := t.pool.BeginTx(ctx, opt)
	if err != nil {
		return nil, err
	}

	if t.cfg.logger != nil {
		return newProxyLogger(t.cfg.logger, &Tx{tx: tx}), nil
	}

	return &Tx{tx: tx}, nil
}

func (t *TxProvider) Transact(ctx context.Context, fn func(tx ConnectionI) error) error {
	tx, err := t.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		tx.Rollback(ctx)
	}()

	err = fn(tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
