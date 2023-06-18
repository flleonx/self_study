package config

import (
	"sync"

	configLogger "auth-server/config/logger"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type ConfigDatabase struct {
	User     string `env:"DB_USER" env-default:"auth_user"`
	Password string `env:"DB_PASSWORD" env-default:"auth123"`
	Port     string `env:"DB_PORT" env-default:"57001"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Name     string `env:"DB_NAME" env-default:"auth"`
	Engine   string `env:"DB_ENGINE" env-default:"postgres"`
	SslMode  string `env:"DB_SSL_MODE" env-default:"disable"`
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
		zap.String("DB_ENGINE", cfgDb.Engine),
		zap.String("DB_USER", cfgDb.User),
		zap.String("DB_PASSWORD", cfgDb.Password),
		zap.String("DB_HOST", cfgDb.Host),
		zap.String("DB_PORT", cfgDb.Port),
		zap.String("DB_NAME", cfgDb.Name),
		zap.String("DB_SSL_MODE", cfgDb.SslMode),
	)

	return cfgDb
}
