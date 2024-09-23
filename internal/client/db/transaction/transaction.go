package transaction

import (
	"context"

	"github.com/pkg/errors"

	"github.com/asp3cto/auth/internal/client/db"
	"github.com/asp3cto/auth/internal/client/db/pg"
	"github.com/jackc/pgx/v4"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager creates a tx manager
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "cant begin transaction")
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrap(err, "errRollback")
			}

			return
		}
		err = tx.Commit(ctx)
		if err != nil {
			err = errors.Wrap(err, "tx commit failed")
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing inside transaction")
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
