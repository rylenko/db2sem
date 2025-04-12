package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"db2sem/internal/config"
	"db2sem/internal/db/pg"
)

func Run(ctx context.Context, cfg config.Config) error {
	_, err := getDBConn(ctx, cfg.PostgresDSN)
	if err != nil {
		return fmt.Errorf("getDBConn(%q): %w", cfg.PostgresDSN, err)
	}

	fmt.Println("connected.")

	return nil
}

func getDBConn(ctx context.Context, dsn string) (pg.Conn, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("New(%q): %w", dsn, err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Ping(): %w", err)
	}

	return pg.NewConn(pool), nil
}
