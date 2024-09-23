package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Client interface for working with db
type Client interface {
	DB() DB
	Close() error
}

// Query is a wrapper for storing query name
type Query struct {
	Name     string
	QueryRaw string
}

// SQLExecer combines NamedExecer and QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer is used fow working with Query in scany
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer is user for working with usual queries
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger interface for checking db availability
type Pinger interface {
	Ping(ctx context.Context) error
}

// Transactor interface defines an entity to do transactions
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Handler is a handler for a transaction
type Handler func(ctx context.Context) error

// TxManager manages transactions
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// DB interface for working with db
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
