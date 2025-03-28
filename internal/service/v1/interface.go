package service_v1

import (
	"context"

	entities_user_v1 "github.com/cliprate/ptfm-auth-svc/internal/entities/user/v1"
)

type UserStoreService interface {
	CreateUser(ctx context.Context, req *entities_user_v1.User_Create) (*entities_user_v1.User_Light, error)
}
