package planetscale

import (
	"context"
	"time"

	"github.com/golerplate/pkg/constants"
	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func (db *dbClient) CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error) {

	userid := constants.GenerateDataPrefixWithULID(constants.User)
	now := time.Now()
	err := db.db.DB.QueryRowContext(ctx,
		`INSERT INTO 
			users 
			(id, username, email, password, is_verified, profile_picture, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);
		`,
		userid, req.Username, req.Email, req.Password, false, "", now, now)
	if err != nil {
		return nil, err.Err()
	}

	return &entities_user_v1.User{
		ID:             userid,
		Username:       req.Username,
		Email:          req.Email,
		IsVerified:     false,
		ProfilePicture: "",
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (d *dbClient) GetUserByEmail(ctx context.Context, email string) (*entities_user_v1.User, error) {
	return &entities_user_v1.User{
		ID:             "1",
		Username:       "testuser",
		Email:          email,
		IsVerified:     false,
		ProfilePicture: "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (d *dbClient) GetUserByID(ctx context.Context, id string) (*entities_user_v1.User, error) {
	return &entities_user_v1.User{
		ID:             id,
		Username:       "testuser",
		Email:          "",
		IsVerified:     false,
		ProfilePicture: "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (d *dbClient) GetUserByUsername(ctx context.Context, username string) (*entities_user_v1.User, error) {
	return &entities_user_v1.User{
		ID:             "1",
		Username:       username,
		Email:          "test",
		IsVerified:     false,
		ProfilePicture: "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}
