package handlers_grpc_user_v1

import (
	"context"
	"errors"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/wrappers"
	userv1 "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
	"github.com/golerplate/pkg/grpc"
	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *handler) CreateUser(ctx context.Context, c *connectgo.Request[userv1.CreateUserRequest]) (*connectgo.Response[userv1.CreateUserResponse], error) {
	if c.Msg.GetUsername() == nil || c.Msg.GetUsername().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid username"))
	}
	if c.Msg.GetEmail() == nil || c.Msg.GetEmail().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid email"))
	}
	if c.Msg.GetPassword() == nil || c.Msg.GetPassword().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid password"))
	}

	user, err := h.userStoreService.CreateUser(ctx, &entities_user_v1.CreateUserRequest{
		Username: c.Msg.GetUsername().GetValue(),
		Email:    c.Msg.GetEmail().GetValue(),
		Password: c.Msg.GetPassword().GetValue(),
	})
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.CreateUserResponse{
		User: &userv1.User{
			Id:               &wrappers.StringValue{Value: user.ID},
			Username:         &wrappers.StringValue{Value: user.Username},
			Email:            &wrappers.StringValue{Value: user.Email},
			IsAdmin:          &wrappers.BoolValue{Value: user.IsAdmin},
			IsBanned:         &wrappers.BoolValue{Value: user.IsBanned},
			HasVerifiedEmail: &wrappers.BoolValue{Value: user.HasVerifiedEmail},
			CreatedAt:        &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:        &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByEmail(ctx context.Context, c *connectgo.Request[userv1.GetUserByEmailRequest]) (*connectgo.Response[userv1.GetUserByEmailResponse], error) {
	if c.Msg.GetEmail() == nil || c.Msg.GetEmail().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid email"))
	}

	user, err := h.userStoreService.GetUserByEmail(ctx, c.Msg.GetEmail().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.GetUserByEmailResponse{
		User: &userv1.User{
			Id:               &wrappers.StringValue{Value: user.ID},
			Username:         &wrappers.StringValue{Value: user.Username},
			Email:            &wrappers.StringValue{Value: user.Email},
			IsAdmin:          &wrappers.BoolValue{Value: user.IsAdmin},
			IsBanned:         &wrappers.BoolValue{Value: user.IsBanned},
			HasVerifiedEmail: &wrappers.BoolValue{Value: user.HasVerifiedEmail},
			CreatedAt:        &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:        &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByID(ctx context.Context, c *connectgo.Request[userv1.GetUserByIDRequest]) (*connectgo.Response[userv1.GetUserByIDResponse], error) {
	if c.Msg.GetId() == nil || c.Msg.GetId().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid id"))
	}

	user, err := h.userStoreService.GetUserByID(ctx, c.Msg.GetId().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.GetUserByIDResponse{
		User: &userv1.User{
			Id:               &wrappers.StringValue{Value: user.ID},
			Username:         &wrappers.StringValue{Value: user.Username},
			Email:            &wrappers.StringValue{Value: user.Email},
			IsAdmin:          &wrappers.BoolValue{Value: user.IsAdmin},
			IsBanned:         &wrappers.BoolValue{Value: user.IsBanned},
			HasVerifiedEmail: &wrappers.BoolValue{Value: user.HasVerifiedEmail},
			CreatedAt:        &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:        &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByUsername(ctx context.Context, c *connectgo.Request[userv1.GetUserByUsernameRequest]) (*connectgo.Response[userv1.GetUserByUsernameResponse], error) {
	if c.Msg.GetUsername() == nil || c.Msg.GetUsername().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid username"))
	}

	user, err := h.userStoreService.GetUserByUsername(ctx, c.Msg.GetUsername().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.GetUserByUsernameResponse{
		User: &userv1.User{
			Id:               &wrappers.StringValue{Value: user.ID},
			Username:         &wrappers.StringValue{Value: user.Username},
			Email:            &wrappers.StringValue{Value: user.Email},
			IsAdmin:          &wrappers.BoolValue{Value: user.IsAdmin},
			IsBanned:         &wrappers.BoolValue{Value: user.IsBanned},
			HasVerifiedEmail: &wrappers.BoolValue{Value: user.HasVerifiedEmail},
			CreatedAt:        &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:        &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) ChangePassword(ctx context.Context, c *connectgo.Request[userv1.ChangePasswordRequest]) (*connectgo.Response[userv1.ChangePasswordResponse], error) {
	if c.Msg.GetUserId() == nil || c.Msg.GetUserId().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid user_id"))
	}
	if c.Msg.GetOldPassword() == nil || c.Msg.GetOldPassword().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid old_password"))
	}
	if c.Msg.GetNewPassword() == nil || c.Msg.GetNewPassword().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid new_password"))
	}

	err := h.userStoreService.ChangePassword(ctx, c.Msg.GetUserId().GetValue(), c.Msg.GetOldPassword().GetValue(), c.Msg.GetNewPassword().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.ChangePasswordResponse{}), nil
}
