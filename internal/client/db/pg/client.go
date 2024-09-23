package pg

import (
	"context"

	"github.com/pkg/errors"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/asp3cto/auth/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

// New creates a db Client
func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %s", err.Error())
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
