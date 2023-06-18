package config

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger
var once sync.Once

func GetLogger() zap.Logger {
	var loadError error
	once.Do(func() {

		newLogger, err := zap.NewDevelopment()

		if err != nil {
			loadError = err
		}

		logger = newLogger

		// defer logger.Sync()
	})

	if loadError != nil {
		log.Fatal("Something happened initializing the logger", loadError)
	}

	return *logger
}
