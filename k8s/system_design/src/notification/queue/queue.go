package queue

import (
	"sync"

	configLogger "notification-server/config/logger"
	configQueue "notification-server/config/queue"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var (
	queueConn *amqp.Connection
	once      sync.Once
)

func Start() *amqp.Connection {
	logger := configLogger.GetLogger()
	var queueError error

	cfgQueue := configQueue.GetQueueConfig()

	once.Do(func() {
		conn, err := amqp.Dial(cfgQueue.Uri)
		if err != nil {
			logger.Panic("Something happened creating queue connection", zap.Error(err))
			queueError = err
			return
		}

		ch, err := conn.Channel()
		if err != nil {
			logger.Panic("Something happened creating queue channel", zap.Error(err))
			queueError = err
			return
		}

		_, err = ch.QueueDeclare(
			cfgQueue.Mp3Queue, // name
			true,              // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // no-wait
			nil,               // arguments
		)

		if err != nil {
			logger.Panic("Error declaring the mp3 queue", zap.Error(err))
			return
		}

		queueConn = conn
	})

	if queueError != nil {
		logger.Fatal("Something happened creating queue connection", zap.Error(queueError))
	}

	return queueConn
}
