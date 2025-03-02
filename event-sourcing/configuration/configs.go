package configuration

import (
	"github.com/caarlos0/env/v11"
)

type (
	Database struct {
		Name     string `env:"DATABASE_NAME"`
		Host     string `env:"DATABASE_HOST"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
	}

	Config struct {
		Database
	}
)

func LoadConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return cfg, err
}
