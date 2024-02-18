package config

import (
	"github.com/Golerplate/pkg/grpc"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	grpc.GRPCServerConfig
	GeneralConfig
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
