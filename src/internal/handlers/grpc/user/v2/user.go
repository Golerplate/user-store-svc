package handlers_grpc_user_v1

import (
	"context"
	"errors"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/wrappers"
	userv2 "github.com/golerplate/contracts/generated/services/user/store/svc/v2"
	"github.com/golerplate/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	entities_user_v2 "github.com/golerplate/user-store-svc/internal/entities/user/v2"
)

func (h *handler) CreateUser(ctx context.Context, c *connectgo.Request[userv2.CreateUserRequest]) (*connectgo.Response[userv2.CreateUserResponse], error) {
	if c.Msg.GetUsername() == nil || c.Msg.GetUsername().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid username"))
	}
	if c.Msg.GetEmail() == nil || c.Msg.GetEmail().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid email"))
	}

	user, err := h.userStoreService.CreateUser(ctx, &entities_user_v2.CreateUserRequest{
		Username: c.Msg.GetUsername().GetValue(),
		Email:    c.Msg.GetEmail().GetValue(),
	})
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv2.CreateUserResponse{
		User: &userv2.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			IsBanned:  &wrappers.BoolValue{Value: user.IsBanned},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByEmail(ctx context.Context, c *connectgo.Request[userv2.GetUserByEmailRequest]) (*connectgo.Response[userv2.GetUserByEmailResponse], error) {
	if c.Msg.GetEmail() == nil || c.Msg.GetEmail().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid email"))
	}

	user, err := h.userStoreService.GetUserByEmail(ctx, c.Msg.GetEmail().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv2.GetUserByEmailResponse{
		User: &userv2.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			IsBanned:  &wrappers.BoolValue{Value: user.IsBanned},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByID(ctx context.Context, c *connectgo.Request[userv2.GetUserByIDRequest]) (*connectgo.Response[userv2.GetUserByIDResponse], error) {
	if c.Msg.GetId() == nil || c.Msg.GetId().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid id"))
	}

	user, err := h.userStoreService.GetUserByID(ctx, c.Msg.GetId().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv2.GetUserByIDResponse{
		User: &userv2.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			IsBanned:  &wrappers.BoolValue{Value: user.IsBanned},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByUsername(ctx context.Context, c *connectgo.Request[userv2.GetUserByUsernameRequest]) (*connectgo.Response[userv2.GetUserByUsernameResponse], error) {
	if c.Msg.GetUsername() == nil || c.Msg.GetUsername().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid username"))
	}

	user, err := h.userStoreService.GetUserByUsername(ctx, c.Msg.GetUsername().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv2.GetUserByUsernameResponse{
		User: &userv2.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			IsBanned:  &wrappers.BoolValue{Value: user.IsBanned},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) UpdateUsername(ctx context.Context, c *connectgo.Request[userv2.UpdateUsernameRequest]) (*connectgo.Response[userv2.UpdateUsernameResponse], error) {
	if c.Msg.GetId() == nil || c.Msg.GetId().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid id"))
	}

	if c.Msg.GetUsername() == nil || c.Msg.GetUsername().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid username"))
	}

	user, err := h.userStoreService.UpdateUsername(ctx, c.Msg.GetId().GetValue(), c.Msg.GetUsername().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv2.UpdateUsernameResponse{
		User: &userv2.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			IsBanned:  &wrappers.BoolValue{Value: user.IsBanned},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}
