package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/golerplate/user-store-svc/internal/config"
	database_connection "github.com/golerplate/user-store-svc/internal/database/connection"
	database_planetscale "github.com/golerplate/user-store-svc/internal/database/pgx"
	handlers_grpc "github.com/golerplate/user-store-svc/internal/handlers/grpc"
	service "github.com/golerplate/user-store-svc/internal/service/v1"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	cfg, err := config.GetServiceConfig()
	if err != nil {
		panic(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	databaseConnection := database_connection.NewDatabaseConnection(ctx, cfg)
	userStoreServiceDatastore := database_planetscale.NewPlanetScaleDatastore(databaseConnection)
	userStoreService, err := service.NewUserStoreService(ctx, userStoreServiceDatastore)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create user store service")
	}

	grpcServer, err := handlers_grpc.NewServer(ctx, cfg.GRPCServerConfig, userStoreService)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create gRPC server")
	}

	if err := grpcServer.Setup(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to setup gRPC server")
	}

	if err := grpcServer.Start(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to start gRPC server")
	}

	<-sigs
	cancel()

	if err := grpcServer.Stop(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to stop gRPC server")
	}

	os.Exit(0)
}
