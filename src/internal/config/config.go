package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/golerplate/pkg/cache/redis"
	"github.com/golerplate/pkg/grpc"
)

type Config struct {
	grpc.GRPCServerConfig
	DatabaseConfig
	redis.RedisConfig
	GeneralConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST"`
	Username string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Port     uint16 `env:"DB_PORT"`
}

type GeneralConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
}

func GetServiceConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
