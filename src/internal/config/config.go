package config

import (
	"github.com/golerplate/pkg/cache/redis"
	"github.com/golerplate/pkg/config"
	database_postgres "github.com/golerplate/pkg/database/postgres"
	"github.com/golerplate/pkg/grpc"
)

type Config struct {
	grpc.GRPCServerConfig
	database_postgres.DatabaseConfig
	redis.RedisConfig
	config.ServiceConfig
}
