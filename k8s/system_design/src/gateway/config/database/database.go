package config

import (
	"sync"

	configLogger "converter-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigDatabase struct {
	Uri string `env:"MONGO_DB_URI" env-default:"mongodb://localhost:27017"`
}

var cfgDb ConfigDatabase
var once sync.Once

func GetDatabaseConfig() ConfigDatabase {
	logger := configLogger.GetLogger()

	var loadError error
	once.Do(func() {
		err := cleanenv.ReadEnv(&cfgDb)
		if err != nil {
			loadError = err
		}
	})

	if loadError != nil {
		logger.Fatal("Something happened initializing database env variables", zap.Error(loadError))
	}

	logger.Debug("Database connection variables",
		zap.String("MONGO_DB_URI", cfgDb.Uri),
	)

	return cfgDb
}
