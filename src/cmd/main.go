package main

import (
	"fmt"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())

	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, syscall.SIGTERM)

	// cfg := config.GetServiceConfig()

	// userStoreServiceDatastore := serviceDatastore.NewPlanetScaleDatastore()
	// userStoreService := service.NewUserStoreService(userStoreServiceDatastore)
	// userStoreServiceHandler := handler_user_v1.NewUserStoreServiceHandler(userStoreService)

	// grpcBuilder := server_user_v1.NewGRPCBuilder(userStoreServiceHandler)

	// grpcServer := http.Server{
	// 	Addr:    fmt.Sprintf(":%d", cfg.GRPCServerConfig.Port),
	// 	Handler: h2c.NewHandler(grpcBuilder, &http2.Server{}),
	// }

	// go func() {
	// 	if err := grpcServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatal().Err(err).Msg("unable to start grpc server")
	// 	}
	// }()

	// <-sigs

	// log.Info().Msg("caught SIGTERM, exiting")
	// cancel()

	// if err := grpcServer.Shutdown(ctx); err != nil {
	// 	log.Error().Err(err).Msg("unable to stop grpc server")
	// }

	// log.Info().Msg("work done; exiting")

	// os.Exit(0)
	fmt.Println("Hello, World!")
}
