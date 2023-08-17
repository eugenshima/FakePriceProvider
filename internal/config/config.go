package config

import (
	"github.com/caarlos0/env/v9"
)

type config struct {
	RedisConnectionString string `env:"REDIS_CONNECTION_STRING"`
}

func NewConfig() (*config, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
