package postgres

import (
	"fmt"
	"context"
	"github.com/dexxp/L0/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect(cfg *config.PG) (*pgxpool.Pool, error) {
	URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	pool, err := pgxpool.Connect(context.Background(), URL)

	return pool, err
}