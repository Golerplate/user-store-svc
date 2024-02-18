package planetscale

import (
	"context"
	"time"

	entities_user_v1 "github.com/Golerplate/user-store-svc/internal/entities/user/v1"
)

func (d *dbClient) CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error) {
	return &entities_user_v1.User{
		ID:        "1",
		Username:  req.Username,
		Email:     req.Email,
		IsAdmin:   false,
		IsBanned:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
