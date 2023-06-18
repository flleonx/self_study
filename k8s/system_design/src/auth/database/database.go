package database

import (
	configDb "auth-server/config/database"
	configLogger "auth-server/config/logger"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func getDatabaseConnString(cfgDatabase configDb.ConfigDatabase) string {

	dbConnString := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfgDatabase.Engine,
		cfgDatabase.User,
		cfgDatabase.Password,
		cfgDatabase.Host,
		cfgDatabase.Port,
		cfgDatabase.Name,
		cfgDatabase.SslMode,
	)

	return dbConnString

}

var db *sql.DB
var once sync.Once

func Start() *sql.DB {
	logger := configLogger.GetLogger()
	var dbError error

	cfgDatabase := configDb.GetDatabaseConfig()

	dbConnString := getDatabaseConnString(cfgDatabase)

	once.Do(func() {
		dbConn, err := sql.Open(cfgDatabase.Engine, dbConnString)

		if err != nil {
			logger.Fatal("Can't connect to database", zap.Error(err))
			dbError = err
			return
		}

		if err := dbConn.Ping(); err != nil {
			logger.Fatal("Unable to reach the database", zap.Error(err))
			dbError = err
			return
		}

		db = dbConn
	})

	if dbError != nil {
		logger.Fatal("Something happened creating db connection", zap.Error(dbError))
	}

	return db
}
