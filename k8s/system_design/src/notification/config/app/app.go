package config

import (
	"sync"

	configLogger "notification-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigApp struct {
	GmailAdress   string `env:"GMAIL_ADDRESS"`
	GmailPassword string `env:"GMAIL_PASSWORD"`
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
