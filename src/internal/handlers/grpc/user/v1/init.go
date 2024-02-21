package handlers_grpc_user_v1

import (
	"context"

	"github.com/golerplate/contracts/generated/services/user/store/svc/v1/svcv1connect"

	service_v1 "github.com/golerplate/user-store-svc/internal/service/v1"
)

type handler struct {
	userStoreService service_v1.UserStoreService
}

func NewUserStoreServiceHandler(ctx context.Context, userStoreService service_v1.UserStoreService) (svcv1connect.UserStoreSvcHandler, error) {
	return &handler{
		userStoreService: userStoreService,
	}, nil
}
