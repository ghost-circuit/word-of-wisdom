package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type key string

const (
	// TxKey is a constant key used to store transaction objects in context.
	TxKey key = "tx"
)

type handler func(ctx context.Context) error

// MakeContextTx returns a new context with the given transaction added to it.
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

// TxManager represents a transaction manager.
type TxManager struct {
	db *ConnectionPool
}

// NewTransactionManager initializes and returns a new transaction TxManager with the given Transactor.
func NewTransactionManager(db *ConnectionPool) *TxManager {
	return &TxManager{
		db: db,
	}
}

// transaction executes the given Handler function within a transaction.
func (m *TxManager) transaction(ctx context.Context, opts pgx.TxOptions, fn handler) (err error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	ctx = MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = fmt.Errorf("err: %w, errRollback: %w", err, errRollback)
			}

			return
		}

		if errCommit := tx.Commit(ctx); errCommit != nil {
			err = fmt.Errorf("tx commit failed: %w", err)
		}
	}()

	if err = fn(ctx); err != nil {
		err = fmt.Errorf("failed executing code inside transaction: %w", err)
	}

	return err
}

// ReadCommitted does a Read Committed transaction.
func (m *TxManager) ReadCommitted(ctx context.Context, f handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
