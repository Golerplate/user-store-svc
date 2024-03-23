package database_pgx_v2

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	database_v2 "github.com/golerplate/user-store-svc/internal/database/v2"
)

type dbClient struct {
	connection *sqlx.DB
}

func NewClient(ctx context.Context, db *sqlx.DB) database_v2.Database {
	return &dbClient{
		connection: db,
	}
}
