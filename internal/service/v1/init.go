package service_v1

import (
	"context"

	database_v1 "github.com/cliprate/ptfm-auth-svc/internal/database/postgres"
)

type service struct {
	store database_v1.Database
}

func NewUserStoreService(ctx context.Context, store database_v1.Database) (*service, error) {
	return &service{
		store: store,
	}, nil
}
