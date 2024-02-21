package database_pgx

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golerplate/pkg/constants"
	"github.com/golerplate/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func (d *dbClient) CreateUser(ctx context.Context, req *entities_user_v1.CreateUserRequest) (*entities_user_v1.User, error) {
	userID := constants.GenerateDataPrefixWithULID(constants.User)
	now := time.Now()

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to hash password: %v", err.Error()))
	}

	_, err = d.connection.DB.ExecContext(ctx,
		`INSERT INTO 
			users (
				id, 
				username,
				email, 
				password, 
				is_admin,
				is_banned,
				has_email_verified, 
				created_at, 
				updated_at
			) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
		`,
		userID, req.Username, req.Email, hashedPassword, false, false, false, now, now)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("failed to create user: %v", err.Error()))
	}

	return &entities_user_v1.User{
		ID:               userID,
		Username:         req.Username,
		Email:            req.Email,
		IsAdmin:          false,
		IsBanned:         false,
		HasVerifiedEmail: false,
		CreatedAt:        now,
		UpdatedAt:        now,
	}, nil
}

func (d *dbClient) GetUserByEmail(ctx context.Context, email string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			username,
			email,
			is_admin,
			is_banned,
			has_email_verified,
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
		&user.IsAdmin,
		&user.IsBanned,
		&user.HasVerifiedEmail,
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
			is_admin,
			is_banned,
			has_email_verified,
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
		&user.IsAdmin,
		&user.IsBanned,
		&user.HasVerifiedEmail,
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
			is_admin,
			is_banned,
			has_email_verified,
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
		&user.IsAdmin,
		&user.IsBanned,
		&user.HasVerifiedEmail,
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

func (d *dbClient) getUserPasswordByID(ctx context.Context, id string) (string, error) {
	var userPassword string

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			password
		FROM
			users
		WHERE
			id = $1;
		`,
		id).Scan(&userPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.NewNotFoundError(fmt.Sprintf("user with id: %s not found", id))
		}

		return "", errors.NewInternalServerError(fmt.Sprintf("failed to get user password by id: %v", err.Error()))
	}

	return userPassword, nil
}

func (d *dbClient) VerifyPassword(ctx context.Context, userID, password string) (bool, error) {
	userPassword, err := d.getUserPasswordByID(ctx, userID)
	if err != nil {
		return false, err
	}

	isPasswordValid, err := checkPasswordHash(password, userPassword)
	if err != nil {
		return false, err
	}

	if !isPasswordValid {
		return false, errors.NewBadRequestError("password is incorrect")
	}

	return true, nil
}

func (d *dbClient) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	userPassword, err := d.getUserPasswordByID(ctx, userID)
	if err != nil {
		return err
	}

	isPasswordValid, err := checkPasswordHash(oldPassword, userPassword)
	if err != nil {
		return err
	}

	if !isPasswordValid {
		return errors.NewBadRequestError("old password is incorrect")
	}

	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	_, err = d.connection.DB.ExecContext(ctx,
		`UPDATE
			users
		SET
			password = $1
		WHERE
			id = $2;
		`,
		hashedPassword, userID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("failed to change password: %v", err.Error()))
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal().Err(err).
			Msg("user: unable to hash password")
		return "", errors.NewInternalServerError(fmt.Sprintf("failed to hash password: %v", err.Error()))
	}
	return string(bytes), err
}

func checkPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Fatal().Err(err).
			Msg("user: unable to hash check password hash")
		return false, errors.NewInternalServerError(fmt.Sprintf("failed to check password hash: %v", err.Error()))
	}

	return true, nil
}
