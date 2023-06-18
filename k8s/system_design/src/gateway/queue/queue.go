package queue

import (
	configLogger "gateway-server/config/logger"
	configQueue "gateway-server/config/queue"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

var queue *amqp.Channel
var once sync.Once

func Start() *amqp.Channel {
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

		queue = ch
	})

	if queueError != nil {
		logger.Fatal("Something happened creating queue interface", zap.Error(queueError))
	}

	return queue
}
