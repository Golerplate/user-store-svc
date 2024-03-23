package database_pgx_v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golerplate/pkg/constants"
	"github.com/golerplate/pkg/errors"
	"github.com/rs/zerolog/log"

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
				created_at, 
				updated_at
			) 
			VALUES ($1, $2, $3, $4, $5);
		`,
		userID, req.Username, req.Email, now, now)
	if err != nil {
		log.Error().Err(err).
			Msgf("failed to create user: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to create user: %v", err.Error()))
	}

	return &entities_user_v1.User{
		ID:        userID,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (d *dbClient) GetUserByEmail(ctx context.Context, email string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			username,
			email,
			created_at,
			updated_at
		FROM
			users
		WHERE
			email = $1
		`,
		email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Msgf("user with email: %s not found", email)
			return nil, errors.NewNotFoundError(fmt.Sprintf("user with email: %s not found", email))
		}

		log.Error().Err(err).
			Msgf("failed to get user by email: %v", err.Error())
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
			created_at, 
			updated_at
		FROM
			users
		WHERE
			id = $1
		`,
		id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Msgf("user with id: %s not found", id)
			return nil, errors.NewNotFoundError(fmt.Sprintf("user with id: %s not found", id))
		}

		log.Error().Err(err).
			Msgf("failed to get user by id: %v", err.Error())
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
			created_at, 
			updated_at
		FROM
			users
		WHERE
			username = $1
		`,
		username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Msgf("user with username: %s not found", username)
			return nil, errors.NewNotFoundError(fmt.Sprintf("user with username: %s not found", username))
		}

		log.Error().Err(err).
			Msgf("failed to get user by username: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to get user by username: %v", err.Error()))
	}

	return user, nil
}

func (d *dbClient) UpdateUsername(ctx context.Context, userID, username string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`UPDATE
			users
		SET
			username = $1,
			updated_at = $2
		WHERE
			id = $3
		RETURNING
			id,
			username,
			email,
			created_at,
			updated_at
		`,
		username, time.Now(), userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Msgf("user: %s not found", userID)
			return nil, errors.NewNotFoundError(fmt.Sprintf("user: %s not found", userID))
		}

		log.Error().Err(err).
			Msgf("failed to update user: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to update user: %v", err.Error()))
	}

	return user, nil
}
