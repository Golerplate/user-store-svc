package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/golerplate/pkg/cache/redis"
	"github.com/golerplate/user-store-svc/internal/config"
	database_pgx "github.com/golerplate/user-store-svc/internal/database/pgx"
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

	cacheConnection := redis.GetConnection(ctx, cfg.RedisConfig)
	cacheRedis := redis.NewRedisCache(ctx, cacheConnection)

	databaseConnection := database_pgx.NewDatabaseConnection(ctx, cfg)

	userStoreService, err := service.NewUserStoreService(ctx, databaseConnection, cacheRedis)
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
