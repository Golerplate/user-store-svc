package service

import (
	"context"

	entity_user_v1 "github.com/Golerplate/user-store-svc/src/internal/entity/user/v1"
)

func (s *service) CreateUser(ctx context.Context, req *entity_user_v1.CreateUserRequest) (*entity_user_v1.User, error) {
	return s.store.CreateUser(ctx, req)
}
