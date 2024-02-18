package handlers_grpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Golerplate/contracts/generated/services/servicesconnect"
	"github.com/Golerplate/contracts/generated/services/user/store/v1/storev1connect"
	"github.com/Golerplate/pkg/grpc"
	pkghandlers "github.com/Golerplate/pkg/grpc/handlers"
	sharedmidlewares "github.com/Golerplate/pkg/grpc/interceptors"
	"github.com/Golerplate/user-store-svc/internal/handlers"
	handlers_grpc_user_v1 "github.com/Golerplate/user-store-svc/internal/handlers/grpc/user/v1"
	"github.com/Golerplate/user-store-svc/internal/service"
	connectgo "github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type grpcServer struct {
	grpcServer *http.Server
	config     grpc.GRPCServerConfig
	service    service.UserStoreService
}

func NewServer(ctx context.Context, cfg grpc.GRPCServerConfig, service service.UserStoreService) (handlers.Server, error) {
	return &grpcServer{
		config:  cfg,
		service: service,
	}, nil
}

func (s *grpcServer) Setup(ctx context.Context) error {
	log.Info().
		Msg("handlers.grpc.grpcServer.Setup: Setting up gRPC server...")

	userStoreServiceHandler := handlers_grpc_user_v1.NewUserStoreServiceHandler(s.service)

	interceptors := connectgo.WithInterceptors(sharedmidlewares.ServerDefaultChain()...)

	reflector := grpcreflect.NewStaticReflector(
		"services.user.store.v1.UserStoreSvc", "services.health.HealthService",
	)

	mux := http.NewServeMux()
	mux.Handle(servicesconnect.NewHealthServiceHandler(pkghandlers.NewHealthHandler()))
	mux.Handle(storev1connect.NewUserStoreSvcHandler(userStoreServiceHandler, interceptors))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	s.grpcServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Port),
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	return nil
}

func (s *grpcServer) Start(ctx context.Context) error {
	log.Info().
		Msg("handlers.grpc.grpcServer.Start: Starting gRPC server...")

	return s.grpcServer.ListenAndServe()
}

func (s *grpcServer) Stop(ctx context.Context) error {
	log.Info().
		Msg("handlers.grpc.grpcServer.Stop: Stopping gRPC server...")

	return s.grpcServer.Shutdown(ctx)
}
