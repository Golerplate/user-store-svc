package planetscale

import (
	"github.com/Golerplate/user-store-svc/internal/datastore"
	_ "github.com/go-sql-driver/mysql"
)

type dbClient struct {
}

func NewPlanetScaleDatastore() datastore.UserStoreServiceDatastore {
	return &dbClient{}
}
