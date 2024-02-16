package handler_user_v1

import (
	"context"
	"errors"

	userv1 "github.com/Golerplate/contracts/generated/services/user/store/v1"
	"github.com/Golerplate/pkg/grpc"
	entity_user_v1 "github.com/Golerplate/user-store-svc/src/internal/entity/user/v1"
	connectgo "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/ptypes/wrappers"
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

	user, err := h.userStoreService.CreateUser(ctx, &entity_user_v1.CreateUserRequest{
		Username: c.Msg.GetUsername().GetValue(),
		Email:    c.Msg.GetEmail().GetValue(),
		Password: c.Msg.GetPassword().GetValue(),
	})
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&userv1.CreateUserResponse{
		User: &userv1.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			IsAdmin:   &wrappers.BoolValue{Value: user.IsAdmin},
			IsBanned:  &wrappers.BoolValue{Value: user.IsBanned},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}
