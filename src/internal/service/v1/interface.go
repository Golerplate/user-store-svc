package service_v1

import (
	"context"

	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

type UserStoreService interface {
	CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities_user_v1.User, error)
	GetUserByID(ctx context.Context, id string) (*entities_user_v1.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities_user_v1.User, error)
	ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error
}
