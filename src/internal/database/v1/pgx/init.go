package database_pgx_v1

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	database_v1 "github.com/golerplate/user-store-svc/internal/database/v1"
)

type dbClient struct {
	connection *sqlx.DB
}

func NewClient(ctx context.Context, db *sqlx.DB) database_v1.Database {
	return &dbClient{
		connection: db,
	}
}
