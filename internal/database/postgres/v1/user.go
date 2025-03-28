package database_v1_pgx

import (
	"context"
	"fmt"
	"time"

	"github.com/cliprate/pkg/constants"
	"github.com/cliprate/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	entities_user_v1 "github.com/cliprate/ptfm-auth-svc/internal/entities/user/v1"
)

func (d *dbClient) CreateUser(ctx context.Context, user *entities_user_v1.User_Create) (*entities_user_v1.User_Light, error) {
	userID := constants.GenerateDataPrefixWithULID(constants.User)
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil, err
	}

	_, err = d.connection.DB.ExecContext(ctx,
		`INSERT INTO 
			users (
				id,
				email,
				username,
				password,
				code,
				created_at,
				updated_at
			) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`,
		userID, user.Email, user.Username, string(hashedPassword), user.Code, now, now)
	if err != nil {
		log.Error().Err(err).
			Msgf("database.postgres.dbClient.CreateUser: failed to create user: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CreateUser: failed to create user: %v", err.Error()))
	}

	return &entities_user_v1.User_Light{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
