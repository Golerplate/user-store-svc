package handler_user_v1

import (
	"github.com/Golerplate/user-store-svc/src/internal/service"
	"github.com/Trade-Arcade-Inc/shared/generated/services/chnl/store/v1/storev1connect"
)

type handler struct {
	userStoreService service.UserStoreService
}

func NewUserStoreServiceHandler(userStoreService service.UserStoreService) storev1connect.ChnlStoreServiceHandler {
	return &handler{
		userStoreService: userStoreService,
	}
}
