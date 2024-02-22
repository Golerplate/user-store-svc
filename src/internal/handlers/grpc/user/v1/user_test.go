package handlers_grpc_user_v1

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	connectgo "github.com/bufbuild/connect-go"
// 	"github.com/golang/protobuf/ptypes/wrappers"
// 	userv1 "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
// 	"github.com/golerplate/pkg/constants"
// 	pkgerrors "github.com/golerplate/pkg/errors"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// 	"google.golang.org/protobuf/types/known/wrapperspb"

// 	database_mocks "github.com/golerplate/user-store-svc/internal/database/mocks"
// 	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
// )

// func Test_CreateUser(t *testing.T) {
// 	t.Run("ok - create user", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		m := database_mocks.NewMockDatabase(ctrl)

// 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// 		created := time.Now()

// 		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
// 			ExternalID: "testuser",
// 			Email:      "testuser@test.com",
// 		}).Return(&entities_user_v1.User{
// 			ID:         userid,
// 			ExternalID: "username",
// 			Username:   "testuser",
// 			Email:      "testuser@test.com",
// 			CreatedAt:  created,
// 			UpdatedAt:  created,
// 		}, nil)

// 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// 		assert.NotNil(t, h)
// 		assert.NoError(t, err)

// 		req := &connectgo.Request[userv1.CreateUserRequest]{
// 			Msg: &userv1.CreateUserRequest{
// 				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
// 				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
// 			},
// 		}

// 		user, err := h.CreateUser(context.Background(), req)
// 		assert.NotNil(t, user)
// 		assert.NoError(t, err)

// 		assert.EqualValues(t, connectgo.NewResponse(&userv1.CreateUserResponse{
// 			User: &userv1.User{
// 				Id:         &wrappers.StringValue{Value: userid},
// 				ExternalId: &wrappers.StringValue{Value: "testuser"},
// 				Username:   &wrappers.StringValue{Value: "testuser"},
// 				Email:      &wrappers.StringValue{Value: "testuser@test.com"},
// 				CreatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// 				UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// 			},
// 		}), user)
// 	})
// 	t.Run("nok - create user", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		m := database_mocks.NewMockDatabase(ctrl)

// 		m.EXPECT().CreateUser(gomock.Any(), &entities_user_v1.CreateUserRequest{
// 			ExternalID: "testuser",
// 			Email:      "testuser@test.com",
// 		}).Return(nil, pkgerrors.NewInternalServerError("error"))

// 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// 		assert.NotNil(t, h)
// 		assert.NoError(t, err)

// 		req := &connectgo.Request[userv1.CreateUserRequest]{
// 			Msg: &userv1.CreateUserRequest{
// 				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
// 				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
// 			},
// 		}

// 		user, err := h.CreateUser(context.Background(), req)
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 	})
// 	t.Run("nok - create user without username", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		m := database_mocks.NewMockDatabase(ctrl)

// 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// 		assert.NotNil(t, h)
// 		assert.NoError(t, err)

// 		req := &connectgo.Request[userv1.CreateUserRequest]{
// 			Msg: &userv1.CreateUserRequest{
// 				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
// 				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
// 			},
// 		}

// 		user, err := h.CreateUser(context.Background(), req)
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
// 	})
// 	t.Run("nok - create user without email", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		m := database_mocks.NewMockDatabase(ctrl)

// 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// 		assert.NotNil(t, h)
// 		assert.NoError(t, err)

// 		req := &connectgo.Request[userv1.CreateUserRequest]{
// 			Msg: &userv1.CreateUserRequest{
// 				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
// 				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
// 			},
// 		}

// 		user, err := h.CreateUser(context.Background(), req)
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
// 	})
// 	t.Run("nok - create user without password", func(t *testing.T) {
// 		ctrl := gomock.NewController(t)
// 		m := database_mocks.NewMockDatabase(ctrl)

// 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// 		assert.NotNil(t, h)
// 		assert.NoError(t, err)

// 		req := &connectgo.Request[userv1.CreateUserRequest]{
// 			Msg: &userv1.CreateUserRequest{
// 				ExternalId: &wrapperspb.StringValue{Value: "testuser"},
// 				Email:      &wrapperspb.StringValue{Value: "testuser@test.com"},
// 			},
// 		}

// 		user, err := h.CreateUser(context.Background(), req)
// 		assert.Nil(t, user)
// 		assert.Error(t, err)
// 		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
// 	})
// }

// // func Test_GetUserByEmail(t *testing.T) {
// // 	t.Run("ok - get user by email", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// // 		created := time.Now()

// // 		m.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(&entities_user_v1.User{
// // 			ID:        userid,
// // 			Username:  "testuser",
// // 			Email:     "testuser@test.com",
// // 			CreatedAt: created,
// // 			UpdatedAt: created,
// // 		}, nil)

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByEmailRequest]{
// // 			Msg: &userv1.GetUserByEmailRequest{
// // 				Email: &wrapperspb.StringValue{Value: "testuser@test.com"},
// // 			},
// // 		}

// // 		user, err := h.GetUserByEmail(context.Background(), req)
// // 		assert.NotNil(t, user)
// // 		assert.NoError(t, err)

// // 		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByEmailResponse{
// // 			User: &userv1.User{
// // 				Id:        &wrappers.StringValue{Value: userid},
// // 				Username:  &wrappers.StringValue{Value: "testuser"},
// // 				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
// // 				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// // 				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// // 			},
// // 		}), user)
// // 	})
// // 	t.Run("nok - get user by email", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		m.EXPECT().GetUserByEmail(gomock.Any(), "testuser@test.com").Return(nil, pkgerrors.NewInternalServerError("error"))

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByEmailRequest]{
// // 			Msg: &userv1.GetUserByEmailRequest{
// // 				Email: &wrapperspb.StringValue{Value: "testuser@test.com"},
// // 			},
// // 		}

// // 		user, err := h.GetUserByEmail(context.Background(), req)
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 	})
// // 	t.Run("nok - get user without email", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByEmailRequest]{
// // 			Msg: &userv1.GetUserByEmailRequest{
// // 				Email: &wrapperspb.StringValue{Value: ""},
// // 			},
// // 		}

// // 		user, err := h.GetUserByEmail(context.Background(), req)
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
// // 	})
// // }

// // func Test_GetUserByID(t *testing.T) {
// // 	t.Run("ok - get user by id", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// // 		created := time.Now()

// // 		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(&entities_user_v1.User{
// // 			ID:        userid,
// // 			Username:  "testuser",
// // 			Email:     "testuser@test.com",
// // 			CreatedAt: created,
// // 			UpdatedAt: created,
// // 		}, nil)

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByIDRequest]{
// // 			Msg: &userv1.GetUserByIDRequest{
// // 				Id: &wrapperspb.StringValue{Value: userid},
// // 			},
// // 		}

// // 		user, err := h.GetUserByID(context.Background(), req)
// // 		assert.NotNil(t, user)
// // 		assert.NoError(t, err)

// // 		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByIDResponse{
// // 			User: &userv1.User{
// // 				Id:        &wrappers.StringValue{Value: userid},
// // 				Username:  &wrappers.StringValue{Value: "testuser"},
// // 				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
// // 				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// // 				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// // 			},
// // 		}), user)
// // 	})
// // 	t.Run("nok - get user by email", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		userid := constants.GenerateDataPrefixWithULID(constants.User)

// // 		m.EXPECT().GetUserByID(gomock.Any(), userid).Return(nil, pkgerrors.NewInternalServerError("error"))

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByIDRequest]{
// // 			Msg: &userv1.GetUserByIDRequest{
// // 				Id: &wrapperspb.StringValue{Value: userid},
// // 			},
// // 		}

// // 		user, err := h.GetUserByID(context.Background(), req)
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 	})
// // 	t.Run("nok - get user without id", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByIDRequest]{
// // 			Msg: &userv1.GetUserByIDRequest{
// // 				Id: &wrapperspb.StringValue{Value: ""},
// // 			},
// // 		}

// // 		user, err := h.GetUserByID(context.Background(), req)
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
// // 	})
// // }

// // func Test_GetUserByUsername(t *testing.T) {
// // 	t.Run("ok - get user by username", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		userid := constants.GenerateDataPrefixWithULID(constants.User)
// // 		created := time.Now()

// // 		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(&entities_user_v1.User{
// // 			ID:        userid,
// // 			Username:  "testuser",
// // 			Email:     "testuser@test.com",
// // 			CreatedAt: created,
// // 			UpdatedAt: created,
// // 		}, nil)

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByUsernameRequest]{
// // 			Msg: &userv1.GetUserByUsernameRequest{
// // 				Username: &wrapperspb.StringValue{Value: "testuser"},
// // 			},
// // 		}

// // 		user, err := h.GetUserByUsername(context.Background(), req)
// // 		assert.NotNil(t, user)
// // 		assert.NoError(t, err)

// // 		assert.EqualValues(t, connectgo.NewResponse(&userv1.GetUserByUsernameResponse{
// // 			User: &userv1.User{
// // 				Id:        &wrappers.StringValue{Value: userid},
// // 				Username:  &wrappers.StringValue{Value: "testuser"},
// // 				Email:     &wrappers.StringValue{Value: "testuser@test.com"},
// // 				CreatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// // 				UpdatedAt: &timestamppb.Timestamp{Seconds: int64(created.Second()), Nanos: int32(created.Nanosecond())},
// // 			},
// // 		}), user)
// // 	})
// // 	t.Run("nok - get user by username", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		m.EXPECT().GetUserByUsername(gomock.Any(), "testuser").Return(nil, pkgerrors.NewInternalServerError("error"))

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByUsernameRequest]{
// // 			Msg: &userv1.GetUserByUsernameRequest{
// // 				Username: &wrapperspb.StringValue{Value: "testuser"},
// // 			},
// // 		}

// // 		user, err := h.GetUserByUsername(context.Background(), req)
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 	})
// // 	t.Run("nok - get user without username", func(t *testing.T) {
// // 		ctrl := gomock.NewController(t)
// // 		m := database_mocks.NewMockDatabase(ctrl)

// // 		h, err := NewUserStoreServiceHandler(context.Background(), m)
// // 		assert.NotNil(t, h)
// // 		assert.NoError(t, err)

// // 		req := &connectgo.Request[userv1.GetUserByUsernameRequest]{
// // 			Msg: &userv1.GetUserByUsernameRequest{
// // 				Username: &wrapperspb.StringValue{Value: ""},
// // 			},
// // 		}

// // 		user, err := h.GetUserByUsername(context.Background(), req)
// // 		assert.Nil(t, user)
// // 		assert.Error(t, err)
// // 		assert.Equal(t, connectgo.CodeInvalidArgument, connectgo.CodeOf(err))
// // 	})
// // }
