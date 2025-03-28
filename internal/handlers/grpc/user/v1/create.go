package handlers_grpc_user_v1

import (
	"context"
	"errors"

	connectgo "github.com/bufbuild/connect-go"
	svcv1connect "github.com/cliprate/contracts/generated/services/ptfm/auth/svc/v1/svcv1connect"
	"github.com/cliprate/pkg/grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"

	entities_user_v1 "github.com/cliprate/ptfm-auth-svc/internal/entities/user/v1"
)

func (h *handler) CreateUser(ctx context.Context, c *connectgo.Request[svcv1connect.CreateUserRequest]) (*connectgo.Response[userv1.CreateUserResponse], error) {
	if c.Msg.GetUsername() == nil || c.Msg.GetUsername().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid username"))
	}

	if c.Msg.GetEmail() == nil || c.Msg.GetEmail().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid email"))
	}

	if c.Msg.GetPassword() == nil || c.Msg.GetPassword().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid password"))
	}

	if c.Msg.GetCode() == nil || c.Msg.GetCode().GetValue() == "" {
		return nil, connectgo.NewError(connectgo.CodeInvalidArgument, errors.New("invalid code"))
	}

	user, err := h.userStoreService.CreateUser(ctx, &entities_user_v1.User_Create{
		Username: c.Msg.GetNickName().GetValue(),
		Password: c.Msg.GetPassword().GetValue(),
		Email:    c.Msg.GetEmail().GetValue(),
		Code:     c.Msg.GetCode().GetValue(),
	})
	if err != nil {
		return nil, grpc.TranslateToGRPCError(ctx, err)
	}

	return connectgo.NewResponse(&svcv1connect.CreateUserResponse{
		User: &svcv1connect.User{
			Id:        &wrappers.StringValue{Value: user.ID},
			Username:  &wrappers.StringValue{Value: user.Username},
			Email:     &wrappers.StringValue{Value: user.Email},
			CreatedAt: &timestamppb.Timestamp{Seconds: int64(user.CreatedAt.Second()), Nanos: int32(user.CreatedAt.Nanosecond())},
			UpdatedAt: &timestamppb.Timestamp{Seconds: int64(user.UpdatedAt.Second()), Nanos: int32(user.UpdatedAt.Nanosecond())},
		},
	}), nil
}
