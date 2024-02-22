package service_v1

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	cache_mocks "github.com/golerplate/pkg/cache/mocks"
// 	"github.com/golerplate/pkg/constants"
// 	pkgerrors "github.com/golerplate/pkg/errors"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"

// 	database_mocks "github.com/golerplate/user-store-svc/internal/database/mocks"
// 	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
// )

// func Test_CreateUser(t *testing.T) {
// 	t.Run("ok - create user", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// 		created := time.Now()

// 		mock_database.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
// 			ExternalID: "testuser",
// 			Email:      "testuser@test.com",
// 		}).Return(&entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, nil)

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.CreateUser(context.Background(), &entities_user_v1.CreateUserRequest{
// 			ExternalID: "testuser",
// 			Email:      "testuser@test.com",
// 		})
// 		assert.NotNil(t, user)
// 		assert.NoError(t, err)

// 		assert.EqualValues(t, &entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, user)
// 	})
// 	t.Run("nok - create user", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		mock_database.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
// 			ExternalID: "testuser",
// 			Email:      "testuser@test.com",
// 		}).Return(nil, pkgerrors.NewInternalServerError("error"))

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.CreateUser(context.Background(), &entities_user_v1.CreateUserRequest{
// 			ExternalID: "testuser",
// 			Email:      "testuser@test.com",
// 		})
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 	})
// }

// func Test_GetUserByEmail(t *testing.T) {
// 	t.Run("ok - get user by email from cache", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// 		created := time.Now()

// 		mock_database.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(&entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, nil)

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return(&entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, nil)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.GetUserByEmail(context.Background(), "testuser@test.com")
// 		assert.NotNil(t, user)
// 		assert.NoError(t, err)

// 		assert.EqualValues(t, &entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, user)
// 	})
// 	t.Run("not found - get user by email", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		mock_database.EXPECT().GetUserByEmail(gomock.Any(), "test@test.com").Return(nil, pkgerrors.NewNotFoundError("user with email: test@test.com not found"))

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.GetUserByEmail(context.Background(), "test@test.com")
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 	})
// 	t.Run("nok - get user by email", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		mock_database.EXPECT().GetUserByEmail(gomock.Any(), "test@test.com").Return(nil, pkgerrors.NewNotFoundError("failed to get user by email: test@test.com"))

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.GetUserByEmail(context.Background(), "test@test.com")
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 	})
// }

// func Test_GetUserByID(t *testing.T) {
// 	t.Run("ok - get user by id", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// 		created := time.Now()

// 		mock_database.EXPECT().GetUserByID(gomock.Any(), userid).Return(&entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, nil)

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.GetUserByID(context.Background(), userid)
// 		assert.NotNil(t, user)
// 		assert.NoError(t, err)

// 		assert.EqualValues(t, &entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "testuser",
// 			Username:   "username",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, user)
// 	})
// 	t.Run("not found - get user by id", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		mock_database.EXPECT().GetUserByID(gomock.Any(), "user_023UN139").Return(nil, pkgerrors.NewNotFoundError("user with id: user_023UN139 not found"))

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.GetUserByID(context.Background(), "user_023UN139")
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 	})
// 	t.Run("nok - get user by id", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		mock_database := database_mocks.NewMockDatabase(ctrl)

// 		mock_database.EXPECT().GetUserByID(gomock.Any(), "user_023UN139").Return(nil, pkgerrors.NewNotFoundError("failed to get user by id: user_023UN139"))

// 		mock_cache := cache_mocks.NewMockCache(ctrl)

// 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// 		assert.NotNil(t, s)
// 		assert.NoError(t, err)

// 		user, err := s.GetUserByID(context.Background(), "user_023UN139")
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 	})
// }

// // func Test_GetUserByUsername(t *testing.T) {
// // 	t.Run("ok - get user by username", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		mock_database := database_mocks.NewMockDatabase(ctrl)

// // 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// // 		created := time.Now()

// // 		mock_database.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(&entities_user_v1.User{
// // 			ID:               userid,
// // 			Username:         "testuser",
// // 			Email:            "testuser@test.com",
// // 			IsAdmin:          false,
// // 			IsBanned:         false,
// // 			HasVerifiedEmail: false,
// // 			CreatedAt:        created,
// // 			UpdatedAt:        created,
// // 		}, nil)

// // 		mock_cache := cache_mocks.NewMockCache(ctrl)

// // 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// // 		assert.NotNil(t, s)
// // 		assert.NoError(t, err)

// // 		user, err := s.GetUserByUsername(context.Background(), "testuser")
// // 		assert.NotNil(t, user)
// // 		assert.NoError(t, err)

// // 		assert.EqualValues(t, &entities_user_v1.User{
// // 			ID:               userid,
// // 			Username:         "testuser",
// // 			Email:            "testuser@test.com",
// // 			IsAdmin:          false,
// // 			IsBanned:         false,
// // 			HasVerifiedEmail: false,
// // 			CreatedAt:        created,
// // 			UpdatedAt:        created,
// // 		}, user)
// // 	})
// // 	t.Run("not found - get user by username", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		mock_database := database_mocks.NewMockDatabase(ctrl)

// // 		mock_database.EXPECT().GetUserByUsername(gomock.Any(), "testtesttest").Return(nil, pkgerrors.NewNotFoundError("user with username: testtesttest not found"))

// // 		mock_cache := cache_mocks.NewMockCache(ctrl)

// // 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// // 		assert.NotNil(t, s)
// // 		assert.NoError(t, err)

// // 		user, err := s.GetUserByUsername(context.Background(), "testtesttest")
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 	})
// // 	t.Run("nok - get user by username", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		mock_database := database_mocks.NewMockDatabase(ctrl)

// // 		mock_database.EXPECT().GetUserByUsername(gomock.Any(), "testtesttest").Return(nil, pkgerrors.NewNotFoundError("failed to get user by username: testtesttest"))

// // 		mock_cache := cache_mocks.NewMockCache(ctrl)

// // 		s, err := NewUserStoreService(context.Background(), mock_database, mock_cache)
// // 		assert.NotNil(t, s)
// // 		assert.NoError(t, err)

// // 		user, err := s.GetUserByUsername(context.Background(), "testtesttest")
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 	})
// // }
