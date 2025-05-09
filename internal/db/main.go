package db

import (
	"context"
	"fmt"

	"github.com/EfosaE/naijacost-api/internal/config"
	"github.com/EfosaE/naijacost-api/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Pool    *pgxpool.Pool  // Connection pool for PostgreSQL (DB)
	Queries *sqlc.Queries // my sqlc package supports pgx instead of database/sql
)

func InitDB() error {
	ctx := context.Background()

	// Create a connection pool instead of a single connection
	var err error
	Pool, err = pgxpool.New(ctx, config.App.Dsn)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Verify the connection
	if err = Pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Initialize queries with the connection pool
	Queries = sqlc.New(Pool)

	return nil
}

// You would also need a cleanup function
func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}
