package service_v1

import (
	"context"
	"testing"
	"time"

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
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
			Username: "testuser",
			Email:    "testuser@test.com",
			Password: "123",
		}).Return(&entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, nil)

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.CreateUser(context.Background(), &entities_user_v1.CreateUserRequest{
			Username: "testuser",
			Email:    "testuser@test.com",
			Password: "123",
		})
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:        userid,
			Username:  "testuser",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}, user)
	})
	t.Run("nok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
			Username: "testuser",
			Email:    "testuser@test.com",
			Password: "123",
		}).Return(nil, pkgerrors.NewInternalServerError("error"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.CreateUser(context.Background(), &entities_user_v1.CreateUserRequest{
			Username: "testuser",
			Email:    "testuser@test.com",
			Password: "123",
		})
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func Test_GetUserByEmail(t *testing.T) {
	t.Run("ok - get user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(&entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, nil)

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "testuser@test.com")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, user)
	})
	t.Run("not found - get user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByEmail(gomock.Any(), "test@test.com").Return(nil, pkgerrors.NewNotFoundError("user with email: test@test.com not found"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "test@test.com")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByEmail(gomock.Any(), "test@test.com").Return(nil, pkgerrors.NewNotFoundError("failed to get user by email: test@test.com"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByEmail(context.Background(), "test@test.com")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func Test_GetUserByID(t *testing.T) {
	t.Run("ok - get user by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(&entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, nil)

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), userid)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, user)
	})
	t.Run("not found - get user by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByID(gomock.Any(), "user_023UN139").Return(nil, pkgerrors.NewNotFoundError("user with id: user_023UN139 not found"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), "user_023UN139")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByID(gomock.Any(), "user_023UN139").Return(nil, pkgerrors.NewNotFoundError("failed to get user by id: user_023UN139"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByID(context.Background(), "user_023UN139")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func Test_GetUserByUsername(t *testing.T) {
	t.Run("ok - get user by username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(&entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, nil)

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "testuser")
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, &entities_user_v1.User{
			ID:             userid,
			Username:       "testuser",
			Email:          "testuser@test.com",
			IsVerified:     false,
			ProfilePicture: "",
			CreatedAt:      created,
			UpdatedAt:      created,
		}, user)
	})
	t.Run("not found - get user by username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByUsername(gomock.Any(), "testtesttest").Return(nil, pkgerrors.NewNotFoundError("user with username: testtesttest not found"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "testtesttest")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user by username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByUsername(gomock.Any(), "testtesttest").Return(nil, pkgerrors.NewNotFoundError("failed to get user by username: testtesttest"))

		s, err := NewUserStoreService(context.Background(), m)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		user, err := s.GetUserByUsername(context.Background(), "testtesttest")
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}
