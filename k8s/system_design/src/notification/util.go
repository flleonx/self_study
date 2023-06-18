package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/smtp"
	"time"

	configApp "notification-server/config/app"
	configLogger "notification-server/config/logger"
	configQueue "notification-server/config/queue"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type QueueMessage struct {
	VideoFId  string `json:"video_fid"`
	Mp3FileId string `json:"mp3_fid"`
	Username  string `json:"username"`
}

func manualNack(ch *amqp.Channel, queueMessage QueueMessage) {
	logger := configLogger.GetLogger()
	cfgQueue := configQueue.GetQueueConfig()
	// NOTE: If something fails I create the message again manually (I still don't know why the channel closes after send ack)
	// NOTE: This is a poor solution
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nackMessage := QueueMessage{
		VideoFId:  queueMessage.VideoFId,
		Mp3FileId: "",
		Username:  queueMessage.Username,
	}

	body, err := json.Marshal(nackMessage)
	if err != nil {
		logger.Error("Error parsing the json nack message body", zap.Error(err))
		return
	}

	err = ch.PublishWithContext(ctx,
		"",                // exchange
		cfgQueue.Mp3Queue, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})

	if err != nil {
		logger.Error("Error sending the nack message to the queue", zap.Error(err))
	}
}

func notification(ch *amqp.Channel, delivery *amqp.Delivery) {
	logger := configLogger.GetLogger()
	cfgApp := configApp.GetAppConfig()

	var queueMessage QueueMessage

	err := json.Unmarshal(delivery.Body, &queueMessage)
	if err != nil {
		logger.Error("Error parsing the queue message", zap.Error(err))
		// WARN: I'm not sure if Nack could close the channel
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}

	logger.Info("Message payload",
		zap.String("video_fid", queueMessage.VideoFId),
		zap.String("mp3_fid", queueMessage.Mp3FileId),
		zap.String("user", queueMessage.Username),
	)

	// Sender data
	from := cfgApp.GmailAdress
	password := cfgApp.GmailPassword

	// Receiver email address
	to := queueMessage.Username

	// smpt server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 465

	// Message
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\nmp3 file_id: %s is now ready!", "MP3 Download", queueMessage.Mp3FileId))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	tlsConfig := &tls.Config{
		ServerName: smtpHost,
	}

	serverName := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	connection, err := tls.Dial("tcp", serverName, tlsConfig)
	if err != nil {
		logger.Error("Error creating connection with the smtp server", zap.Error(err))
		// manualNack(ch, queueMessage)
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}

	smtpClient, err := smtp.NewClient(connection, smtpHost)

	if err = smtpClient.Auth(auth); err != nil {
		logger.Error("Error setting auth", zap.Error(err))
		// manualNack(ch, queueMessage)
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}
	if err = smtpClient.Mail(from); err != nil {
		logger.Error("Error setting from", zap.Error(err))
		// manualNack(ch, queueMessage)
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}

	if err = smtpClient.Rcpt(to); err != nil {
		logger.Error("Error setting receptors", zap.Error(err))
		// manualNack(ch, queueMessage)
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}

	writer, err := smtpClient.Data()
	if err != nil {
		logger.Error("Error creating writer", zap.Error(err))
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}

	// Sending email
	_, err = writer.Write([]byte(message))
	if err != nil {
		logger.Error("Error sending email", zap.Error(err))
		ch.Nack(delivery.DeliveryTag, false, true)
		return
	}

	// Close writer
	if err := writer.Close(); err != nil {
		logger.Error("Error closing writer", zap.Error(err))
		return
	}

	// Close tcp connection
	if err := smtpClient.Quit(); err != nil {
		logger.Error("Error closing connection", zap.Error(err))
		return
	}

	logger.Info("Mail sent")
	return
}
