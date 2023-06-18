package main

import (
	"context"
	"time"

	configLogger "notification-server/config/logger"
	configQueue "notification-server/config/queue"
	"notification-server/queue"

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
		cfgQueue.Mp3Queue, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		logger.Fatal("failed to register the consumer", zap.Error(err))
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			logger.Info("A message have been recevied")
			notification(ch, &d)
			logger.Info("notification finished")
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
