package planetscale

import (
	"context"

	entity_user_v1 "github.com/Golerplate/user-store-svc/src/internal/entity/user/v1"
)

func (d *dbClient) CreateUser(ctx context.Context, req *entity_user_v1.CreateUserRequest) (*entity_user_v1.User, error) {
	return nil, nil
}
