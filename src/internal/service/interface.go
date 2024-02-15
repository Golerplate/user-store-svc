package service

import (
	"context"

	entity_user_v1 "github.com/Golerplate/user-store-svc/src/internal/entity/user/v1"
)

type UserStoreService interface {
	CreateUser(ctx context.Context, req *entity_user_v1.CreateUserRequest) (*entity_user_v1.User, error)
}
