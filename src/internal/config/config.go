package config

import (
	"github.com/Golerplate/pkg/grpc"
	"github.com/caarlos0/env/v8"
	"github.com/rs/zerolog/log"
)

type Config struct {
	grpc.GRPCServerConfig
	GeneralConfig
}

type GeneralConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
}

func GetServiceConfig() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err).Msg("unable to build config")
	}

	return &cfg
}
