package handlers_grpc_user_v1

import (
	"context"

	"github.com/golerplate/contracts/generated/services/user/store/svc/v2/svcv2connect"

	service_v2 "github.com/golerplate/user-store-svc/internal/service/v2"
)

type handler struct {
	userStoreService service_v2.UserStoreService
}

func NewUserStoreServiceHandler(ctx context.Context, userStoreService service_v2.UserStoreService) (svcv2connect.UserStoreSvcHandler, error) {
	return &handler{
		userStoreService: userStoreService,
	}, nil
}
