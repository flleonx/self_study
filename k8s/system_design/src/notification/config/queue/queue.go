package config

import (
	"sync"

	configLogger "notification-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigQueue struct {
	Uri        string `env:"QUEUE_URI" env-default:"amqp://guest:guest@rabbitmq:5672"`
	VideoQueue string `env:"VIDEO_QUEUE" env-default:"video"`
	Mp3Queue   string `env:"MP3_QUEUE" env-default:"mp3"`
}

var (
	cfgQueue ConfigQueue
	once     sync.Once
)

func GetQueueConfig() ConfigQueue {
	var loadError error
	logger := configLogger.GetLogger()

	once.Do(func() {
		err := cleanenv.ReadEnv(&cfgQueue)
		if err != nil {
			loadError = err
		}
	})

	if loadError != nil {
		logger.Fatal("Something happened initializing queue env variables", zap.Error(loadError))
	}

	return cfgQueue
}
