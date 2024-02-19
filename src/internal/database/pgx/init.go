package database_pgx

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/golerplate/user-store-svc/internal/config"
	"github.com/golerplate/user-store-svc/internal/database"
)

type dbClient struct {
	connection *sqlx.DB
}

func NewDatabaseConnection(ctx context.Context, cfg *config.Config) database.Database {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Username,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.DBName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create user store db")
	}

	return &dbClient{
		connection: db,
	}
}
