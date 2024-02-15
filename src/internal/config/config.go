package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/rs/zerolog/log"
)

type Config struct {
	HTTPServerConfig
	GRPCServerConfig
	GeneralConfig
}

type GeneralConfig struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
}

type HTTPServerConfig struct {
	Port uint16 `env:"HTTP_SERVER_PORT" envDefault:"3003"`
}

type GRPCServerConfig struct {
	Port uint16 `env:"GRPC_SERVER_PORT" envDefault:"50051"`
}

func GetServiceConfig() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err).Msg("unable to build config")
	}

	return &cfg
}
