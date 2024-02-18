package handlers_grpc_user_v1

import (
	"github.com/Golerplate/contracts/generated/services/user/store/v1/storev1connect"
	"github.com/Golerplate/user-store-svc/internal/service"
)

type handler struct {
	userStoreService service.UserStoreService
}

func NewUserStoreServiceHandler(userStoreService service.UserStoreService) storev1connect.UserStoreSvcHandler {
	return &handler{
		userStoreService: userStoreService,
	}
}
