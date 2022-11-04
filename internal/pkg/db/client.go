package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close() error
	DB() *DB
}

type client struct {
	db        *DB
	closeFunc context.CancelFunc
}

func NewClient(ctx context.Context, connString string) (Client, error) {
	dbc, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithCancel(ctx)

	return &client{
		db:        &DB{pool: dbc},
		closeFunc: cancel,
	}, nil
}

func (c *client) Close() error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}
		if c.db != nil {
			c.db.pool.Close()
		}
	}

	return nil
}

func (c *client) DB() *DB {
	return c.db
}
