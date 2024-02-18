package planetscale

import (
	"github.com/jmoiron/sqlx"

	"github.com/golerplate/user-store-svc/internal/database"
)

type dbClient struct {
	db *sqlx.DB
}

func NewPlanetScaleDatastore(db *sqlx.DB) database.Database {
	return &dbClient{
		db: db,
	}
}
