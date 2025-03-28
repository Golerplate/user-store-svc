package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	pkg_config "github.com/cliprate/pkg/config"
	pkg_postgres "github.com/cliprate/pkg/database/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/cliprate/ptfm-auth-svc/internal/config"
	database_v1_pgx "github.com/cliprate/ptfm-auth-svc/internal/database/postgres/v1"
	handlers_grpc "github.com/cliprate/ptfm-auth-svc/internal/handlers/grpc"
	service "github.com/cliprate/ptfm-auth-svc/internal/service/v1"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	cfg := &config.Config{}
	err := pkg_config.ParseConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to parse config")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	databaseConnection, err := pkg_postgres.NewDatabaseConnection(ctx, &cfg.DatabaseConfig)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create database connection")
	}
	databaseClient := database_v1_pgx.NewClient(ctx, databaseConnection)

	userStoreService, err := service.NewUserStoreService(ctx, databaseClient)
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
