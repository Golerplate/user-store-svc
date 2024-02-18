package datastore

import (
	"context"

	entities_user_v1 "github.com/Golerplate/user-store-svc/internal/entities/user/v1"
)

type UserStoreServiceDatastore interface {
	CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error)
}
