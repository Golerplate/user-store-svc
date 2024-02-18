package handlers_grpc_user_v1

import (
	"github.com/Golerplate/contracts/generated/services/user/store/svc/v1/svcv1connect"
	service_v1 "github.com/Golerplate/user-store-svc/internal/service/v1"
)

type handler struct {
	userStoreService service_v1.UserStoreService
}

func NewUserStoreServiceHandler(userStoreService service_v1.UserStoreService) svcv1connect.UserStoreSvcHandler {
	return &handler{
		userStoreService: userStoreService,
	}
}
