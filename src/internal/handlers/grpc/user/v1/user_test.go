package handlers_grpc_user_v1

import (
	"context"
	"testing"
	"time"

	user_store_svc_v1_entities "github.com/Golerplate/contracts/clients/user-store-svc/v1/entities"
	user_store_svc_v1_mocks "github.com/Golerplate/contracts/clients/user-store-svc/v1/mocks"
	"github.com/Golerplate/pkg/constants"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	entities_user_v1 "github.com/Golerplate/user-store-svc/internal/entities/user/v1"
	service_v1 "github.com/Golerplate/user-store-svc/internal/service/v1"
)

func Test_CreateUser(t *testing.T) {
	t.Run("ok - username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := user_store_svc_v1_mocks.NewMockUserStoreSvc(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
			Username: "testuser",
			Email:    "testuser@test.com",
			Password: "123",
		}).Return(&user_store_svc_v1_entities.User{
			ID:        userid,
			Username:  "testuser",
			Email:     "testuser@test.com",
			IsAdmin:   false,
			IsBanned:  false,
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		s, err := service_v1.NewUserStoreService(context.Background(), m)
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
}
