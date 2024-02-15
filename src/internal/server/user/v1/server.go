package server_user_v1

import (
	"net/http"

	"github.com/Trade-Arcade-Inc/shared/generated/services/chnl/store/v1/storev1connect"
	"github.com/Trade-Arcade-Inc/shared/generated/services/servicesconnect"
	"github.com/Trade-Arcade-Inc/shared/pkg/grpc/handlers"
	connectgo "github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	connect_go_prometheus "github.com/easyCZ/connect-go-prometheus"
)

func NewGRPCBuilder(
	chnlStoreServiceHandler storev1connect.ChnlStoreServiceHandler,
) *http.ServeMux {
	interceptorchain := append(sharedmidlewares.ServerDefaultChain(), connect_go_prometheus.NewInterceptor())
	interceptors := connectgo.WithInterceptors(interceptorchain...)

	reflector := grpcreflect.NewStaticReflector(
		"services.user.store.v1.UserStoreService", "services.health.HealthService",
	)

	mux := http.NewServeMux()
	mux.Handle(servicesconnect.NewHealthServiceHandler(handlers.NewHealthHandler()))
	mux.Handle(storev1connect.NewChnlStoreServiceHandler(chnlStoreServiceHandler, interceptors))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return mux
}
