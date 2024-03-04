package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func (s *service) clearCacheForUser(ctx context.Context, user *entities_user_v1.User) error {
	err := s.cache.Del(ctx, user.ID)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", user.ID).
			Msg("service.v1.service.clearCacheForUser: unable to delete user from cache by user_id")
		return err
	}

	err = s.cache.Del(ctx, user.Username)
	if err != nil {
		log.Error().Err(err).
			Str("username", user.Username).
			Msg("service.v1.service.clearCacheForUser: unable to delete user from cache by username")
		return err
	}

	err = s.cache.Del(ctx, user.Email)
	if err != nil {
		log.Error().Err(err).
			Str("email", user.Email).
			Msg("service.v1.service.clearCacheForUser: unable to delete user from cache by email")
		return err
	}

	return nil
}

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

func (s *service) UpdateUsername(ctx context.Context, userID, username string) (*entities_user_v1.User, error) {
	user, err := s.store.UpdateUsername(ctx, userID, username)
	if err != nil {
		return nil, err
	}

	err = s.clearCacheForUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
