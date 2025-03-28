package config

import (
	pkg_config "github.com/cliprate/pkg/config"
	database_postgres "github.com/cliprate/pkg/database/postgres"
	"github.com/cliprate/pkg/grpc"
)

type Config struct {
	ServiceConfig    pkg_config.ServiceConfig
	GRPCServerConfig grpc.GRPCServerConfig
	DatabaseConfig   database_postgres.PostgresConfig
}
