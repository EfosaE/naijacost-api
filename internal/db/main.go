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

type DB struct {
    Pool    *pgxpool.Pool
    Queries *sqlc.Queries
}

func InitDB(ctx context.Context) (*DB, error) {
    // Create a connection pool
    pool, err := pgxpool.New(ctx, config.App.Dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }

    // Verify the connection
    if err = pool.Ping(ctx); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }
    
    fmt.Printf("Connection to database successful ✅\n")

    // Initialize queries with the connection pool
    queries := sqlc.New(pool)

    return &DB{
        Pool:    pool,
        Queries: queries,
    }, nil
}

