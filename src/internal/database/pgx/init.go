package database_pgx

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golerplate/user-store-svc/internal/database"
)

type dbClient struct {
	connection *sqlx.DB
}

func NewClient(ctx context.Context, db *sqlx.DB) database.Database {
	return &dbClient{
		connection: db,
	}
}
