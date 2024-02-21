package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

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
	cachedUser, err := s.cache.Get(ctx, email)
	if err == nil {
		var user *entities_user_v1.User
		err = json.Unmarshal([]byte(cachedUser), &user)
		if err != nil {
			log.Error().Err(err).
				Str("email", email).
				Msg("service.v1.service.GetUserByEmail: unable to unmarshall user")
		} else {
			return user, nil
		}
	}

	user, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Error().Err(err).
			Str("email", email).
			Msg("service.v1.service.GetUserByEmail: unable to marshal user")
	} else {
		_ = s.cache.SetEx(ctx, generateUserCacheKeyWithEmail(email), bytes, userCacheDuration)
	}

	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, userID string) (*entities_user_v1.User, error) {
	cachedUser, err := s.cache.Get(ctx, userID)
	if err == nil {
		var user *entities_user_v1.User
		err = json.Unmarshal([]byte(cachedUser), &user)
		if err != nil {
			log.Error().Err(err).
				Str("user_id", userID).
				Msg("service.v1.service.GetUserByID: unable to unmarshall user")
		} else {
			return user, nil
		}
	}

	user, err := s.store.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Msg("service.v1.service.GetUserByID: unable to marshal user")
	} else {
		_ = s.cache.SetEx(ctx, generateUserCacheKeyWithUserID(userID), bytes, userCacheDuration)
	}

	return user, nil
}

func (s *service) GetUserByUsername(ctx context.Context, username string) (*entities_user_v1.User, error) {
	cachedUser, err := s.cache.Get(ctx, username)
	if err == nil {
		var user *entities_user_v1.User
		err = json.Unmarshal([]byte(cachedUser), &user)
		if err != nil {
			log.Error().Err(err).
				Str("username", username).
				Msg("service.v1.service.GetUserByID: unable to unmarshall user")
		} else {
			return user, nil
		}
	}

	user, err := s.store.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Error().Err(err).
			Str("username", username).
			Msg("service.v1.service.GetUserByID: unable to marshal user")
	} else {
		_ = s.cache.SetEx(ctx, generateUserCacheKeyWithUsername(username), bytes, userCacheDuration)
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
