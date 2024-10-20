package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DatabaseClient represents a PostgresQL database client.
type DatabaseClient struct {
	masterDBC *ConnectionPool
}

// NewClient initializes a new PostgreSQL client with the given data source name (DatabaseDSN).
func NewClient(ctx context.Context, dsn string) (*DatabaseClient, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &DatabaseClient{
		masterDBC: &ConnectionPool{dbc: dbc},
	}, nil
}

// DB returns the master database connection.
func (c *DatabaseClient) DB() *ConnectionPool {
	return c.masterDBC
}

// Close closes the master database connection.
func (c *DatabaseClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.dbc.Close()
	}

	return nil
}

// Ping checks the connection to the database.
func (c *DatabaseClient) Ping(ctx context.Context) error {
	return c.masterDBC.dbc.Ping(ctx)
}
