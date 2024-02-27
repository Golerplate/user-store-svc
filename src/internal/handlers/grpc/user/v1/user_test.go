package handlers_grpc_user_v1

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/wrappers"
	userv1 "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
	cache_mocks "github.com/golerplate/pkg/cache/mocks"
	"github.com/golerplate/pkg/constants"
	pkgerrors "github.com/golerplate/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	database_mocks "github.com/golerplate/user-store-svc/internal/database/mocks"
	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
	service_v1 "github.com/golerplate/user-store-svc/internal/service/v1"
)

func Test_CreateUser(t *testing.T) {
	t.Run("ok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		userid := constants.GenerateDataPrefixWithULID(constants.User)
		created := time.Now()

		m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.CreateUserRequest]{
			Msg: &userv1.CreateUserRequest{
				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.CreateUser(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv1.CreateUserResponse{
			User: &userv1.User{
				Id:         &wrappers.StringValue{Value: userid},
				ExternalId: &wrappers.StringValue{Value: "testuser"},
				Username:   &wrappers.StringValue{Value: gomock.Any().String()},
				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - create user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Do(
			func(ctx context.Context, req *entities_user_v1.ServiceCreateUserRequest) {
				assert.Equal(t, "testuser", req.ExternalID)
				assert.Equal(t, "testuser@test.com", req.Email)
				assert.NotEmpty(t, req.Username)
			},
		).Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.CreateUserRequest]{
			Msg: &userv1.CreateUserRequest{
				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.CreateUser(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - create user without external_id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.CreateUserRequest]{
			Msg: &userv1.CreateUserRequest{
				ExternalId: &wrapperspb.StringValue{Value: ""},
				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.CreateUserRequest]{
			Msg: &userv1.CreateUserRequest{
				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
				Email:      &wrapperspb.StringValue{Value: ""},
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByEmailRequest]{
			Msg: &userv1.GetUserByEmailRequest{
				Email: &wrapperspb.StringValue{Value: "testuser@test.com"},
			},
		}

		user, err := h.GetUserByEmail(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByEmailResponse{
			User: &userv1.User{
				Id:         &wrappers.StringValue{Value: userid},
				ExternalId: &wrappers.StringValue{Value: "testuser"},
				Username:   &wrappers.StringValue{Value: "username"},
				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByEmailRequest]{
			Msg: &userv1.GetUserByEmailRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByEmailRequest]{
			Msg: &userv1.GetUserByEmailRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByIDRequest]{
			Msg: &userv1.GetUserByIDRequest{
				Id: &wrapperspb.StringValue{Value: userid},
			},
		}

		user, err := h.GetUserByID(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByIDResponse{
			User: &userv1.User{
				Id:         &wrappers.StringValue{Value: userid},
				ExternalId: &wrappers.StringValue{Value: "testuser"},
				Username:   &wrappers.StringValue{Value: "username"},
				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByIDRequest]{
			Msg: &userv1.GetUserByIDRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByIDRequest]{
			Msg: &userv1.GetUserByIDRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByUsernameRequest]{
			Msg: &userv1.GetUserByUsernameRequest{
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.GetUserByUsername(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByUsernameResponse{
			User: &userv1.User{
				Id:         &wrappers.StringValue{Value: userid},
				ExternalId: &wrappers.StringValue{Value: "testuser"},
				Username:   &wrappers.StringValue{Value: "username"},
				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByUsernameRequest]{
			Msg: &userv1.GetUserByUsernameRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByUsernameRequest]{
			Msg: &userv1.GetUserByUsernameRequest{
				Username: &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.GetUserByUsername(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
	})
}

func Test_GetUserByExternalID(t *testing.T) {
	t.Run("ok - get user by external_id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByExternalIDRequest]{
			Msg: &userv1.GetUserByExternalIDRequest{
				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
			},
		}

		user, err := h.GetUserByExternalID(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByExternalIDResponse{
			User: &userv1.User{
				Id:         &wrappers.StringValue{Value: userid},
				ExternalId: &wrappers.StringValue{Value: "testuser"},
				Username:   &wrappers.StringValue{Value: "username"},
				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - get user by external_id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		m.EXPECT().GetUserByExternalID(gomock.Any(), "testuser").Return(nil, pkgerrors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "testuser").Return(fakeData, nil)

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByExternalIDRequest]{
			Msg: &userv1.GetUserByExternalIDRequest{
				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
			},
		}

		user, err := h.GetUserByExternalID(context.Background(), req)
		assert.Nil(t, user)
		assert.Error(t, err)
	})
	t.Run("nok - get user without external_id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.GetUserByExternalIDRequest]{
			Msg: &userv1.GetUserByExternalIDRequest{
				ExternalId: &wrapperspb.StringValue{Value: ""},
			},
		}

		user, err := h.GetUserByExternalID(context.Background(), req)
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

		m.EXPECT().UpdateUsername(gomock.Any(), userid, "username").Return(&entities_user_v1.User{
			ID:         userid,
			ExternalID: "testuser",
			Username:   gomock.Any().String(),
			Email:      "testuser@test.com",
			CreatedAt:  created,
			UpdatedAt:  created,
		}, nil)

		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)
		mock_cache.EXPECT().Del(gomock.Any(), gomock.Any()).Return(nil)

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.UpdateUsernameRequest]{
			Msg: &userv1.UpdateUsernameRequest{
				Id:       &wrapperspb.StringValue{Value: userid},
				Username: &wrapperspb.StringValue{Value: "username"},
			},
		}

		user, err := h.UpdateUsername(context.Background(), req)
		assert.NotNil(t, user)
		assert.NoError(t, err)

		assert.EqualValues(t, connectgo.NewResponse(&userv1.UpdateUsernameResponse{
			User: &userv1.User{
				Id:         &wrappers.StringValue{Value: userid},
				ExternalId: &wrappers.StringValue{Value: "testuser"},
				Username:   &wrappers.StringValue{Value: gomock.Any().String()},
				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
			},
		}), user)
	})
	t.Run("nok - update username", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		m := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		m.EXPECT().UpdateUsername(gomock.Any(), "user_id", "username").Return(nil, pkgerrors.NewInternalServerError("error"))

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.UpdateUsernameRequest]{
			Msg: &userv1.UpdateUsernameRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.UpdateUsernameRequest]{
			Msg: &userv1.UpdateUsernameRequest{
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

		service, err := service_v1.NewUserStoreService(context.Background(), m, mock_cache)
		assert.NotNil(t, service)
		assert.NoError(t, err)

		h, err := NewUserStoreServiceHandler(context.Background(), service)
		assert.NotNil(t, h)
		assert.NoError(t, err)

		req := &connectgo.Request[userv1.UpdateUsernameRequest]{
			Msg: &userv1.UpdateUsernameRequest{
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
