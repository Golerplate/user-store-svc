package service_v1

import (
	"context"

	"github.com/Golerplate/user-store-svc/internal/datastore"
)

type service struct {
	store datastore.UserStoreServiceDatastore
}

func NewUserStoreService(ctx context.Context, store datastore.UserStoreServiceDatastore) (*service, error) {
	return &service{
		store: store,
	}, nil
}
