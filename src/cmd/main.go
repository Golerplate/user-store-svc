package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Golerplate/user-store-svc/src/internal/config"
	handler_user_v1 "github.com/Golerplate/user-store-svc/src/internal/handler/user/v1"
	server_user_v1 "github.com/Golerplate/user-store-svc/src/internal/server/user/v1"
	"github.com/Trade-Arcade-Inc/shared/pkg/http/server"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Golerplate/user-store-svc/src/internal/service"
	serviceDatastore "github.com/Golerplate/user-store-svc/src/internal/service/datastore/planetscale"
)

func main() {
	_, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	cfg := config.GetServiceConfig()

	userStoreServiceDatastore := serviceDatastore.NewPlanetScaleDatastore()
	userStoreService := service.NewUserStoreService(userStoreServiceDatastore)
	userStoreServiceHandler := handler_user_v1.NewUserStoreServiceHandler(userStoreService)

	grpcBuilder := server_user_v1.NewGRPCBuilder(userStoreServiceHandler)

	grpcServer := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.GRPCServerConfig.Port),
		Handler: h2c.NewHandler(grpcBuilder, &http2.Server{}),
	}

	httpServer := server.InitializeTelemetryServer(cfg.HTTPServerConfig.Port)

	go func() {
		if err := grpcServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("unable to start grpc server")
		}
	}()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("unable to start http server")
		}
	}()

	<-sigs

	log.Info().Msg("caught SIGTERM, exiting")
	cancel()

	println("Hello, World!!")
}
