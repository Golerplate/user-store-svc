package service_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	cache_mocks "github.com/golerplate/pkg/cache/mocks"
	"github.com/golerplate/pkg/constants"
	pkgerrors "github.com/golerplate/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	database_mocks "github.com/golerplate/user-store-svc/internal/database/mocks"
	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func Test_CreateUser(t *testing.T) {
	t.Run("ok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(
			func(ctx context.Context, req *entities_user_v1.ServiceCreateUserRequest) {
				assert.Equal(t, "testuser", req.ExternalID)
				assert.Equal(t, "testuser@test.com", req.Email)
				assert.NotEmpty(t, req.Username)
			},
		).Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   gomock.Any().String(),
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.CreateUser(context.Background(), &entities_user_v1.GRPCCreateUserRequest{
			ExternalID: "testuser",
			Email:      "testuser@test.com",
		})
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, gomock.Any().String(), user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("nok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_database.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(
			func(ctx context.Context, req *entities_user_v1.ServiceCreateUserRequest) {
				assert.Equal(t, "testuser", req.ExternalID)
				assert.Equal(t, "testuser@test.com", req.Email)
				assert.NotEmpty(t, req.Username)
			},
		).Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.CreateUser(context.Background(), &entities_user_v1.GRPCCreateUserRequest{
			ExternalID: "testuser",
			Email:      "testuser@test.com",
		})
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func Test_GetUserByEmail(t *testing.T) {
	t.Run("ok - get user by email from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return(string(userCachedBytes), nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "testuser@test.com")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("ok - get user by email from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return("", pkgerrors.NewNotFoundError("error"))

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), "user-store-svc:user:email:testuser@test.com", userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "testuser@test.com")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("nok - get user by email from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_database.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(nil, pkgerrors.NewNotFoundError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return("", pkgerrors.NewNotFoundError("error"))

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "testuser@test.com")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("ok - get user by email when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return(fakeData, nil)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), "user-store-svc:user:email:testuser@test.com", userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "testuser@test.com")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
}

func Test_GetUserByUsename(t *testing.T) {
	t.Run("ok - get user by username from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), "username").Return(string(userCachedBytes), nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "username")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("ok - get user by username from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().GetUserByUsername(gomock.Any(), "username").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "username").Return("", pkgerrors.NewNotFoundError("error"))

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), "user-store-svc:user:username:username", userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "username")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("nok - get user by username from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_database.EXPECT().GetUserByUsername(gomock.Any(), "username").Return(nil, pkgerrors.NewNotFoundError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "username").Return("", pkgerrors.NewNotFoundError("error"))

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "username")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("ok - get user by username when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "username").Return(fakeData, nil)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().GetUserByUsername(gomock.Any(), "username").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), "user-store-svc:user:username:username", userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "username")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
}

func Test_GetUserByID(t *testing.T) {
	t.Run("ok - get user by id from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), userid).Return(string(userCachedBytes), nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), userid)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("ok - get user by id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().GetUserByID(gomock.Any(), userid).Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), userid).Return("", pkgerrors.NewNotFoundError("error"))

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("user-store-svc:user:user_id:%+v", userid), userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), userid)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("nok - get user by id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)

		mock_database.EXPECT().GetUserByID(gomock.Any(), userid).Return(nil, pkgerrors.NewNotFoundError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), userid).Return("", pkgerrors.NewNotFoundError("error"))

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), userid)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("ok - get user by id when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), userid).Return(fakeData, nil)

		mock_database.EXPECT().GetUserByID(gomock.Any(), userid).Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("user-store-svc:user:user_id:%+v", userid), userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), userid)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
}

func Test_GetUserByExternalID(t *testing.T) {
	t.Run("ok - get user by external_id from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser").Return(string(userCachedBytes), nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByExternalID(context.Background(), "testuser")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("ok - get user by external_id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_database.EXPECT().GetUserByExternalID(gomock.Any(), "testuser").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser").Return("", pkgerrors.NewNotFoundError("error"))

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("user-store-svc:user:external_id:%+v", "testuser"), userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByExternalID(context.Background(), "testuser")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("nok - get user by external_id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_database.EXPECT().GetUserByID(gomock.Any(), "testuser").Return(nil, pkgerrors.NewNotFoundError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser").Return("", pkgerrors.NewNotFoundError("error"))

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), "testuser")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("ok - get user by external_id when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "testuser").Return(fakeData, nil)

		mock_database.EXPECT().GetUserByExternalID(gomock.Any(), "testuser").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		userCached := &entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("user-store-svc:user:external_id:%+v", "testuser"), userCachedBytes, time.Hour*24).Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByExternalID(context.Background(), "testuser")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
}

func Test_UpdateUsername(t *testing.T) {
	t.Run("ok - update username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)
		mock_database.EXPECT().UpdateUsername(gomock.Any(), userid, "username").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   "username",
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache.EXPECT().Del(gomock.Any(), userid).Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), "testuser").Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), "username").Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), "testuser@test.com").Return(nil)

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.UpdateUsername(context.Background(), userid, "username")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.Equal(t, userid, user.ID)
		assert.Equal(t, "testuser", user.ExternalID)
		assert.Equal(t, "username", user.Username)
		assert.Equal(t, "testuser@test.com", user.Email)
		assert.True(t, user.CreatedAt.Equal(created))
		assert.True(t, user.UpdatedAt.Equal(created))
	})
	t.Run("nok - update username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)

		mock_cache := cache_mocks.NewMockCache(ctrl)
		mock_database.EXPECT().UpdateUsername(gomock.Any(), userid, "username").Return(nil, pkgerrors.NewInternalServerError("error"))

		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.UpdateUsername(context.Background(), userid, "username")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}
