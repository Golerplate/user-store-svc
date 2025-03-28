package service_v1

import (
	"context"
	"fmt"

	"github.com/cliprate/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	entities_user_v1 "github.com/cliprate/ptfm-auth-svc/internal/entities/user/v1"
)

func (s *service) CreateUser(ctx context.Context, req *entities_user_v1.User_Create) (*entities_user_v1.User_Light, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).
			Msgf("service.v1.service.CreateUser: failed to hash password: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("service.v1.service.CreateUser: failed to hash password: %v", err.Error()))
	}

	user, err := s.store.CreateUser(ctx, &entities_user_v1.User_Create{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
