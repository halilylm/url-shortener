package config

import (
	"github.com/caarlos0/env"
	"time"
)

type Config struct {
	Host                    string        `env:"HOST"`
	Port                    int           `env:"PORT"`
	WriteTimeout            time.Duration `env:"WRITE_TIMEOUT"`
	ReadTimeout             time.Duration `env:"READ_TIMEOUT"`
	IdleTimeout             time.Duration `env:"IDLE_TIMEOUT"`
	ReadHeaderTimeout       time.Duration `env:"READ_HEADER_TIMEOUT"`
	GracefulShutdownTimeout time.Duration `env:"GRACEFUL_SHUTDOWN_TIMEOUT"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
