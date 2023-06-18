package main

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"time"

	configLogger "gateway-server/config/logger"
	"gateway-server/database"
	"gateway-server/queue"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.uber.org/zap"
)

func upload(file *multipart.FileHeader, access claimResp) error {
	logger := configLogger.GetLogger()
	mongoClient := database.Start()
	db := mongoClient.Database("videos")
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		logger.Error("Error creating a new GridFS bucket", zap.Error(err))
		return err
	}

	fileStream, err := file.Open()
	if err != nil {
		logger.Error("Error opening the file attached in the request", zap.Error(err))
		return err
	}

	fileId, err := bucket.UploadFromStream(file.Filename, io.Reader(fileStream))
	if err != nil {
		logger.Error("Error uploading the file to mongoDB", zap.Error(err))
		return err
	}

	type QueueMessage struct {
		VideoFId  string `json:"video_fid"`
		Mp3FileId string `json:"mp3_fid"`
		Username  string `json:"username"`
	}

	hexId := fileId.String()[10:34]

	queueMessage := QueueMessage{
		VideoFId:  hexId,
		Mp3FileId: "",
		Username:  access.Email,
	}

	ch := queue.Start()

	q, err := ch.QueueDeclare(
		"video", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		logger.Error("Error declaring the video queue", zap.Error(err))
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(queueMessage)
	if err != nil {
		logger.Error("Error parsing the json queue message body", zap.Error(err))
		return err
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})

	if err != nil {
		logger.Error("Error sending the message to the queue", zap.Error(err))
		if err := bucket.Delete(fileId); err != nil {
			logger.Error("Error trying to delete file in mongoDB", zap.Error(err))
			return err
		}
	}

	return nil
}
