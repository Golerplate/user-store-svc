package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Golerplate/user-store-svc/internal/config"
	handlers_grpc "github.com/Golerplate/user-store-svc/internal/handlers/grpc"
	"github.com/Golerplate/user-store-svc/internal/service"
	serviceDatastore "github.com/Golerplate/user-store-svc/internal/service/datastore/planetscale"
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

	userStoreServiceDatastore := serviceDatastore.NewPlanetScaleDatastore()
	userStoreService := service.NewUserStoreService(userStoreServiceDatastore)

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
