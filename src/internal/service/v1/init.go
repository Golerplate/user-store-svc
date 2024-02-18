package service_v1

import (
	"context"

	"github.com/golerplate/user-store-svc/internal/database"
)

type service struct {
	store database.Database
}

func NewUserStoreService(ctx context.Context, store database.Database) (*service, error) {
	return &service{
		store: store,
	}, nil
}
