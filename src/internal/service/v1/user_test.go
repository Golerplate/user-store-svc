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
