package handlers_grpc_user_v1

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/wrappers"
	userv2 "github.com/golerplate/contracts/generated/services/user/store/svc/v2"
	cache_mocks "github.com/golerplate/pkg/cache/mocks"
	"github.com/golerplate/pkg/constants"
	pkgerrors "github.com/golerplate/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	database_mocks "github.com/golerplate/user-store-svc/internal/database/v2/mocks"
	entities_user_v2 "github.com/golerplate/user-store-svc/internal/entities/user/v2"
	service_v2 "github.com/golerplate/user-store-svc/internal/service/v2"
)

func Test_CreateUser(t *testing.T) {
	t.Run("ok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v2.CreateUserRequest{
			Username: "Teyz",
			Email:    "testuser@test.com",
		}).Return(&entities_user_v2.User{
			ID:        userid,
			Username:  "Teyz",
			Email:     "testuser@test.com",
			IsBanned:  false,
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.CreateUserRequest]{
			Msg: &userv2.CreateUserRequest{
				Username: &wrapperspb.StringValue{Value: "Teyz"},
				Email:    &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.CreateUser(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv2.CreateUserResponse{
			User: &userv2.User{
				Id:        &wrappers.StringValue{Value: userid},
				Username:  &wrappers.StringValue{Value: "Teyz"},
				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v2.CreateUserRequest{
			Username: "Teyz",
			Email:    "testuser@test.com",
		}).Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.CreateUserRequest]{
			Msg: &userv2.CreateUserRequest{
				Username: &wrapperspb.StringValue{Value: "Teyz"},
				Email:    &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.CreateUser(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - create user without username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.CreateUserRequest]{
			Msg: &userv2.CreateUserRequest{
				Username: &wrapperspb.StringValue{Value: ""},
				Email:    &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.CreateUser(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
	t.Run("nok - create user without email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.CreateUserRequest]{
			Msg: &userv2.CreateUserRequest{
				Username: &wrapperspb.StringValue{Value: "testuser"},
				Email:    &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.CreateUser(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
}

func Test_GetUserByEmail(t *testing.T) {
	t.Run("ok - get user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v2.User{
			ID:        userid,
			Username:  "username",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return(string(userCachedBytes), nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByEmailRequest]{
			Msg: &userv2.GetUserByEmailRequest{
				Email: &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.GetUserByEmail(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv2.GetUserByEmailResponse{
			User: &userv2.User{
				Id:        &wrappers.StringValue{Value: userid},
				Username:  &wrappers.StringValue{Value: "username"},
				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - get user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "testuser@test.com").Return(fakeData, nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByEmailRequest]{
			Msg: &userv2.GetUserByEmailRequest{
				Email: &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.GetUserByEmail(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user without email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByEmailRequest]{
			Msg: &userv2.GetUserByEmailRequest{
				Email: &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.GetUserByEmail(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
}

func Test_GetUserByID(t *testing.T) {
	t.Run("ok - get user by id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v2.User{
			ID:        userid,
			Username:  "username",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), userid).Return(string(userCachedBytes), nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByIDRequest]{
			Msg: &userv2.GetUserByIDRequest{
				Id: &wrapperspb.StringValue{Value: userid},
			},
		}

		user, err := h.GetUserByID(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv2.GetUserByIDResponse{
			User: &userv2.User{
				Id:        &wrappers.StringValue{Value: userid},
				Username:  &wrappers.StringValue{Value: "username"},
				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - get user by email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)

		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), userid).Return(fakeData, nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByIDRequest]{
			Msg: &userv2.GetUserByIDRequest{
				Id: &wrapperspb.StringValue{Value: userid},
			},
		}

		user, err := h.GetUserByID(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user without id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByIDRequest]{
			Msg: &userv2.GetUserByIDRequest{
				Id: &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.GetUserByID(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
}

func Test_GetUserByUsername(t *testing.T) {
	t.Run("ok - get user by username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		userCached := &entities_user_v2.User{
			ID:        userid,
			Username:  "username",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}

		userCachedBytes, _ := json.Marshal(userCached)

		mock_cache.EXPECT().Get(gomock.Any(), "username").Return(string(userCachedBytes), nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByUsernameRequest]{
			Msg: &userv2.GetUserByUsernameRequest{
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.GetUserByUsername(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv2.GetUserByUsernameResponse{
			User: &userv2.User{
				Id:        &wrappers.StringValue{Value: userid},
				Username:  &wrappers.StringValue{Value: "username"},
				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - get user by username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByUsername(gomock.Any(), "username").Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "username").Return(fakeData, nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByUsernameRequest]{
			Msg: &userv2.GetUserByUsernameRequest{
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.GetUserByUsername(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user without username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.GetUserByUsernameRequest]{
			Msg: &userv2.GetUserByUsernameRequest{
				Username: &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.GetUserByUsername(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
}

func Test_UpdateUsername(t *testing.T) {
	t.Run("ok - update username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		m.EXPECT().UpdateUsername(gomock.Any(), userid, "username").Return(&entities_user_v2.User{
			ID:        userid,
			Username:  "username",
			Email:     "testuser@test.com",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.UpdateUsernameRequest]{
			Msg: &userv2.UpdateUsernameRequest{
				Id:       &wrapperspb.StringValue{Value: userid},
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.UpdateUsername(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv2.UpdateUsernameResponse{
			User: &userv2.User{
				Id:        &wrappers.StringValue{Value: userid},
				Username:  &wrappers.StringValue{Value: "username"},
				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - update username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		m.EXPECT().UpdateUsername(gomock.Any(), "user_id", "username").Return(nil, pkgerrors.NewInternalServerError("error"))

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.UpdateUsernameRequest]{
			Msg: &userv2.UpdateUsernameRequest{
				Id:       &wrapperspb.StringValue{Value: "user_id"},
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.UpdateUsername(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - update username without user_id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.UpdateUsernameRequest]{
			Msg: &userv2.UpdateUsernameRequest{
				Id:       &wrapperspb.StringValue{Value: ""},
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.UpdateUsername(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user without username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v2.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv2.UpdateUsernameRequest]{
			Msg: &userv2.UpdateUsernameRequest{
				Id:       &wrapperspb.StringValue{Value: "user_id"},
				Username: &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.UpdateUsername(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
}
