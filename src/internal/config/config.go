package config

import (
	"github.com/golerplate/pkg/cache/redis"
	"github.com/golerplate/pkg/config"
	"github.com/golerplate/pkg/database"
	"github.com/golerplate/pkg/grpc"
)

type Config struct {
	grpc.GRPCServerConfig
	database.DatabaseConfig
	redis.RedisConfig
	config.ServiceConfig
}
