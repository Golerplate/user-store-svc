package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/golerplate/pkg/grpc"
)

type Config struct {
	grpc.GRPCServerConfig
	DatabaseConfig
	GeneralConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Username string `env:"DB_USER" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"root"`
	DBName   string `env:"DB_NAME" envDefault:"user-store-db"`
	Port     uint16 `env:"DB_PORT" envDefault:"5432"`
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
