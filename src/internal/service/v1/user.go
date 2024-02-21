package service_v1

import (
	"context"

	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func (s *service) CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error) {
	user, err := s.store.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*entities_user_v1.User, error) {
	user, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*entities_user_v1.User, error) {
	user, err := s.store.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByUsername(ctx context.Context, username string) (*entities_user_v1.User, error) {
	user, err := s.store.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	err := s.store.ChangePassword(ctx, userID, oldPassword, newPassword)
	if err != nil {
		return err
	}

	return nil
}
