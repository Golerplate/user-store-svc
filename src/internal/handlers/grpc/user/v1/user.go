package handlers_grpc_user_v1

import (
	"context"
	"errors"

	connectgo "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/wrappers"
	userv1 "github.com/golerplate/contracts/generated/services/user/store/svc/v1"
	"github.com/golerplate/pkg/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	entities_user_v1 "github.com/golerplate/user-store-svc/internal/entities/user/v1"
)

func (h *handler) CreateUser(ctx context.Context, c *connectgo.Request[userv1.CreateUserRequest]) (*connectgo.Response[userv1.CreateUserResponse], error) {
	if c.Msg.GetExternalId() == nil || c.Msg.GetExternalId().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid external_id"))
	}
	if c.Msg.GetEmail() == nil || c.Msg.GetEmail().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid email"))
	}

	log.Info().Msg("CreateUser")

	user, err := h.userStoreService.CreateUser(ctx, &entities_user_v1.GRPCCreateUserRequest{
		ExternalID: c.Msg.GetExternalId().GetValue(),
		Email:      c.Msg.GetEmail().GetValue(),
	})
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.CreateUserResponse{
		User: &userv1.User{
			Id:         &wrappers.StringValue{Value: user.ID},
			ExternalId: &wrappers.StringValue{Value: user.ExternalID},
			Username:   &wrappers.StringValue{Value: user.Username},
			Email:      &wrappers.StringValue{Value: user.Email},
			CreatedAt:  &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
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
			Id:         &wrappers.StringValue{Value: user.ID},
			ExternalId: &wrappers.StringValue{Value: user.ExternalID},
			Username:   &wrappers.StringValue{Value: user.Username},
			Email:      &wrappers.StringValue{Value: user.Email},
			CreatedAt:  &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
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
			Id:         &wrappers.StringValue{Value: user.ID},
			ExternalId: &wrappers.StringValue{Value: user.ExternalID},
			Username:   &wrappers.StringValue{Value: user.Username},
			Email:      &wrappers.StringValue{Value: user.Email},
			CreatedAt:  &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
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
			Id:         &wrappers.StringValue{Value: user.ID},
			ExternalId: &wrappers.StringValue{Value: user.ExternalID},
			Username:   &wrappers.StringValue{Value: user.Username},
			Email:      &wrappers.StringValue{Value: user.Email},
			CreatedAt:  &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) GetUserByExternalID(ctx context.Context, c *connectgo.Request[userv1.GetUserByExternalIDRequest]) (*connectgo.Response[userv1.GetUserByExternalIDResponse], error) {
	if c.Msg.GetExternalId() == nil || c.Msg.GetExternalId().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid external_id"))
	}

	user, err := h.userStoreService.GetUserByExternalID(ctx, c.Msg.GetExternalId().GetValue())
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.GetUserByExternalIDResponse{
		User: &userv1.User{
			Id:         &wrappers.StringValue{Value: user.ID},
			ExternalId: &wrappers.StringValue{Value: user.ExternalID},
			Username:   &wrappers.StringValue{Value: user.Username},
			Email:      &wrappers.StringValue{Value: user.Email},
			CreatedAt:  &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}

func (h *handler) UpdateUsername(ctx context.Context, c *connectgo.Request[userv1.UpdateUsernameRequest]) (*connectgo.Response[userv1.UpdateUsernameResponse], error) {
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

	return connectgo.NewResponse(&userv1.UpdateUsernameResponse{
		User: &userv1.User{
			Id:         &wrappers.StringValue{Value: user.ID},
			ExternalId: &wrappers.StringValue{Value: user.ExternalID},
			Username:   &wrappers.StringValue{Value: user.Username},
			Email:      &wrappers.StringValue{Value: user.Email},
			CreatedAt:  &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt:  &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}
