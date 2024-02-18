package planetscale

import (
	"context"

	entities_user_v1 "github.com/Golerplate/user-store-svc/internal/entities/user/v1"
)

func (d *dbClient) CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error) {
	return nil, nil
}
