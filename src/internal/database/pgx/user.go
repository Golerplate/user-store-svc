package database_pgx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golerplate/pkg/constants"
	"github.com/golerplate/pkg/errors"
	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func (d *dbClient) CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error) {
	userID := constants.GenerateDataPrefixWithULID(constants.User)
	now := time.Now()
	_, err := d.connection.DB.ExecContext(ctx,
		`INSERT INTO 
			users (
				id, 
				username,
				email, 
				password, 
				is_verified, 
				profile_picture, 
				created_at, 
				updated_at
			) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
		`,
		userID, req.Username, req.Email, req.Password, false, "", now, now)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to create user: %v", err.Error()))
	}

	return &entities_user_v1.User{
		ID:             userID,
		Username:       req.Username,
		Email:          req.Email,
		IsVerified:     false,
		ProfilePicture: "",
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (d *dbClient) GetUserByEmail(ctx context.Context, email string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			username,
			email,
			is_verified,
			profile_picture,
			created_at,
			updated_at
		FROM
			users
		WHERE
			email = $1;
		`,
		email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.IsVerified,
		&user.ProfilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(fmt.Sprintf("user with email: %s not found", email))
		}

		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to get user by email: %v", err.Error()))
	}

	return user, nil
}

func (d *dbClient) GetUserByID(ctx context.Context, id string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			username,
			email,
			is_verified,
			profile_picture,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = $1;
		`,
		id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.IsVerified,
		&user.ProfilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(fmt.Sprintf("user with id: %s not found", id))
		}

		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to get user by id: %v", err.Error()))
	}

	return user, nil
}

func (d *dbClient) GetUserByUsername(ctx context.Context, username string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			username,
			email,
			is_verified,
			profile_picture,
			created_at,
			updated_at
		FROM
			users
		WHERE
			username = $1;
		`,
		username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.IsVerified,
		&user.ProfilePicture,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(fmt.Sprintf("user with username: %s not found", username))
		}

		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to get user by username: %v", err.Error()))
	}

	return user, nil
}
