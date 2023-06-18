package config

import (
	"sync"

	configLogger "gateway-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigQueue struct {
	Uri string `env:"QUEUE_URI" env-default:"amqp://guest:guest@rabbitmq:5672"`
}

var cfgQueue ConfigQueue
var once sync.Once

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
