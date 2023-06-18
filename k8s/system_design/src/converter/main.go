package main

import (
	"context"
	"time"

	configLogger "converter-server/config/logger"
	configQueue "converter-server/config/queue"
	"converter-server/queue"

	// amqp "github.com/rabbitmq/amqp091-go"

	"go.uber.org/zap"
)

func main() {
	cfgQueue := configQueue.GetQueueConfig()
	logger := configLogger.GetLogger()
	queueConn := queue.Start()
	ch, err := queueConn.Channel()
	if err != nil {
		logger.Fatal("Error creating queue channel", zap.Error(err))
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		cfgQueue.VideoQueue, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	if err != nil {
		logger.Fatal("failed to register the consumer", zap.Error(err))
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			logger.Info("A message have been recevied")
			callback(ch, queueConn, &d)
			logger.Info("Callback finished")
		}
	}()

	logger.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer logger.Sync()
	// TODO: Implement graceful shutdown
	// os.Exit(0)
	return
}
