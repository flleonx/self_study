package config

import (
	"sync"

	configLogger "converter-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigApp struct {
	VideoQueue         string `env:"VIDEO_QUEUE" env-default:"video"`
	Mp3Queue           string `env:"MP3_QUEUE" env-default:"mp3"`
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
