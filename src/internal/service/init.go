package service

import (
	"github.com/Golerplate/user-store-svc/src/internal/service/datastore"
)

type service struct {
	store datastore.UserStoreServiceDatastore
}

func NewUserStoreService(store datastore.UserStoreServiceDatastore) UserStoreService {
	return &service{
		store: store,
	}
}