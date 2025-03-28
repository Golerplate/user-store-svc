package database_v1

import (
	"context"

	entities_user_v1 "github.com/cliprate/ptfm-auth-svc/internal/entities/user/v1"
)

//go:generate mockgen -source interface.go -destination mocks/mock_database.go -package database_mocks
type Database interface {
	//User
	CreateUser(ctx context.Context, req *entities_user_v1.User_Create) (*entities_user_v1.User_Light, error)
}
