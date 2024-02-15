package planetscale

import (
	"github.com/Golerplate/user-store-svc/src/internal/service/datastore"
	_ "github.com/go-sql-driver/mysql"
)

type dbClient struct {
}

func NewPlanetScaleDatastore() datastore.UserStoreServiceDatastore {
	return &dbClient{}
}
