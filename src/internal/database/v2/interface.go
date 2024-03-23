package database_v2

import (
	"context"

	entities_user_v2 "github.com/golerplate/user-store-svc/internal/entities/user/v2"
)

//go:generate mockgen -source interface.go -destination mocks/mock_database.go -package database_mocks
type Database interface {
	CreateUser(ctx context.Context, req *entities_user_v2.CreateUserRequest) (*entities_user_v2.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities_user_v2.User, error)
	GetUserByID(ctx context.Context, id string) (*entities_user_v2.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities_user_v2.User, error)

	UpdateUsername(ctx context.Context, userID, username string) (*entities_user_v2.User, error)
}
