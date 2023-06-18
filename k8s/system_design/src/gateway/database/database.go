package database

import (
	"context"
	configDb "gateway-server/config/database"
	configLogger "gateway-server/config/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func getDatabaseConnString(cfgDatabase configDb.ConfigDatabase) string {

	dbConnString := cfgDatabase.Uri

	return dbConnString

}

var db *mongo.Client
var once sync.Once

func Start() *mongo.Client {
	logger := configLogger.GetLogger()
	var dbError error

	cfgDatabase := configDb.GetDatabaseConfig()

	dbConnString := getDatabaseConnString(cfgDatabase)

	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnString))

		if err != nil {
			logger.Fatal("Error creating the connection", zap.Error(err))
			dbError = err
			return
		}

		// defer func() {
		// 	if err = client.Disconnect(ctx); err != nil {
		// 		logger.Panic("Something happened creating db connection", zap.Error(err))
		// 		dbError = err
		// 	}
		// }()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			logger.Fatal("Unable to reach the database", zap.Error(err))
			dbError = err
			return
		}

		db = client
	})

	if dbError != nil {
		logger.Fatal("Something happened creating db connection", zap.Error(dbError))
	}

	return db
}
