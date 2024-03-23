package service_v1

import (
	"context"

	entities_user_v2 "github.com/golerplate/user-store-svc/internal/entities/user/v2"
)

type UserStoreService interface {
	CreateUser(ctx context.Context, req *entities_user_v2.CreateUserRequest) (*entities_user_v2.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities_user_v2.User, error)
	GetUserByID(ctx context.Context, id string) (*entities_user_v2.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities_user_v2.User, error)

	UpdateUsername(ctx context.Context, userID, username string) (*entities_user_v2.User, error)
}
