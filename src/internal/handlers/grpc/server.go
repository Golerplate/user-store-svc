package handlers_grpc

import (
	"context"
	"fmt"
	"net/http"

	connectgo "github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/golerplate/contracts/generated/services"
	"github.com/golerplate/contracts/generated/services/servicesconnect"
	"github.com/golerplate/contracts/generated/services/user/store/svc/v1/svcv1connect"
	"github.com/golerplate/pkg/grpc"
	sharedmidlewares "github.com/golerplate/pkg/grpc/interceptors"
	"github.com/golerplate/user-store-svc/internal/handlers"
	handlers_grpc_user_v1 "github.com/golerplate/user-store-svc/internal/handlers/grpc/user/v1"
	service "github.com/golerplate/user-store-svc/internal/service/v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type health struct{}

func (h *health) CheckHealth(ctx context.Context, c *connectgo.Request[services.HealthRequest]) (*connectgo.Response[services.HealthResponse], error) {
	return connectgo.NewResponse(&services.HealthResponse{}), nil
}

func NewHealthHandler() servicesconnect.HealthServiceHandler {
	return &health{}
}

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

	userStoreServiceHandler, err := handlers_grpc_user_v1.NewUserStoreServiceHandler(ctx, s.service)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create user store service handler")
	}

	interceptors := connectgo.WithInterceptors(sharedmidlewares.ServerDefaultChain()...)

	reflector := grpcreflect.NewStaticReflector(
		"services.user.store.svc.v1.UserStoreSvc", "services.health.HealthService",
	)

	mux := http.NewServeMux()
	mux.Handle(servicesconnect.NewHealthServiceHandler(NewHealthHandler()))
	mux.Handle(svcv1connect.NewUserStoreSvcHandler(userStoreServiceHandler, interceptors))
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
