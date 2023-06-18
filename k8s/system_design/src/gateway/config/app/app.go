package config

import (
	"sync"

	configLogger "gateway-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigApp struct {
	Host               string `env:"APP_HOST" env-default:"localhost"`
	Port               string `env:"APP_PORT" env-default:"8080"`
	AuthServiceAddress string `env:"APP_AUTH_SERVICE_ADRRESS" env-default:"localhost:5000"`
}

var (
	cfgApp ConfigApp
	once   sync.Once
)

func GetAppConfig() ConfigApp {
	var loadError error
	logger := configLogger.GetLogger()

	once.Do(func() {
		err := cleanenv.ReadEnv(&cfgApp)
		if err != nil {
			loadError = err
		}
	})

	if loadError != nil {
		logger.Fatal("Something happened initializing database env variables", zap.Error(loadError))
	}

	return cfgApp
}
