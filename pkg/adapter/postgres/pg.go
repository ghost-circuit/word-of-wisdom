package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectionPool represents a connection pool to a PostgresQL database.
type ConnectionPool struct {
	dbc *pgxpool.Pool
}

// NewDB initializes a new ConnectionPool instance with the given connection pool.
func NewDB(dbc *pgxpool.Pool) *ConnectionPool {
	return &ConnectionPool{
		dbc: dbc,
	}
}

// ScanOne executes a query and scans the result into the provided destination.
func (p *ConnectionPool) ScanOne(ctx context.Context, dest any, q Query, args ...any) error {
	row, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAll executes a query and scans all the results into the provided destination.
func (p *ConnectionPool) ScanAll(ctx context.Context, dest any, q Query, args ...any) error {
	rows, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// Exec executes a query without returning any rows.
func (p *ConnectionPool) Exec(ctx context.Context, q Query, args ...any) (pgconn.CommandTag, error) {
	logQuery(q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// Query executes a query and returns the resulting rows.
func (p *ConnectionPool) Query(ctx context.Context, q Query, args ...any) (pgx.Rows, error) {
	logQuery(q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRow executes a query and returns a single row.
func (p *ConnectionPool) QueryRow(ctx context.Context, q Query, args ...any) pgx.Row {
	logQuery(q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx begins a new transaction with the given options.
func (p *ConnectionPool) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}
