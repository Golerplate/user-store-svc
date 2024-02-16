package server_user_v1

import (
	"net/http"

	"github.com/Golerplate/contracts/generated/services/servicesconnect"
	"github.com/Golerplate/contracts/generated/services/user/store/v1/storev1connect"
	handlers "github.com/Golerplate/pkg/grpc/handlers"
	sharedmidlewares "github.com/Golerplate/pkg/grpc/interceptors"
	connectgo "github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
)

func NewGRPCBuilder(
	userStoreServiceHandler storev1connect.UserStoreServiceHandler,
) *http.ServeMux {
	interceptors := connectgo.WithInterceptors(sharedmidlewares.ServerDefaultChain()...)

	reflector := grpcreflect.NewStaticReflector(
		"services.user.store.v1.UserStoreSvc", "services.health.HealthService",
	)

	mux := http.NewServeMux()
	mux.Handle(servicesconnect.NewHealthServiceHandler(handlers.NewHealthHandler()))
	mux.Handle(storev1connect.NewUserStoreServiceHandler(userStoreServiceHandler, interceptors))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return mux
}
