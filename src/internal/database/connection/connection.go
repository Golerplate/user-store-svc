package database_connection

import (
	"context"
	"fmt"

	"github.com/golerplate/user-store-svc/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewDatabaseConnection(ctx context.Context, cfg *config.Config) *sqlx.DB {
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

	return db
}
