package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	configLogger "converter-server/config/logger"
	configQueue "converter-server/config/queue"
	"converter-server/database"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.uber.org/zap"
)

type claimResp struct {
	Email string `json:"email"`
	Admin bool   `json:"admin"`
}

type QueueMessage struct {
	VideoFId  string `json:"video_fid"`
	Mp3FileId string `json:"mp3_fid"`
	Username  string `json:"username"`
}

func callback(videoCh *amqp.Channel, queueConn *amqp.Connection, delivery *amqp.Delivery) {
	logger := configLogger.GetLogger()
	cfgQueue := configQueue.GetQueueConfig()

	mongoClient := database.Start()

	videosDb := mongoClient.Database("videos")
	videosBucket, err := gridfs.NewBucket(videosDb)
	if err != nil {
		logger.Error("Error creating a bucket for videos database", zap.Error(err))
		videoCh.Nack(delivery.DeliveryTag, false, true)
		return
	}

	mp3sDb := mongoClient.Database("mp3s")
	mp3sBucket, err := gridfs.NewBucket(mp3sDb)
	if err != nil {
		logger.Error("Error creating a bucket for mp3s database", zap.Error(err))
		videoCh.Nack(delivery.DeliveryTag, false, true)
		return
	}

	var queueMessage QueueMessage

	err = json.Unmarshal(delivery.Body, &queueMessage)

	if err != nil {
		logger.Error("Error parsing message queue payload", zap.Error(err))
		videoCh.Nack(delivery.DeliveryTag, false, true)
		return
	}

	logger.Info("Message payload",
		zap.String("video_fid", queueMessage.VideoFId),
		zap.String("mp3_fid", queueMessage.Mp3FileId),
		zap.String("user", queueMessage.Username),
	)

	// Publish channel
	// TODO: find out if is optimal create new channel each time
	ch, err := queueConn.Channel()
	if err != nil {
		logger.Error("Error trying to create publish channel", zap.Error(err))
	}
	defer ch.Close()

	err = start(queueMessage, videosBucket, mp3sBucket, ch)

	// NOTE: Why the channel closes after send the ack????
	// if err != nil {
	// 	logger.Error("Error - preparing Nack", zap.Error(err))
	// 	err = videoCh.Nack(delivery.DeliveryTag, false, true)
	// 	if err != nil {
	// 		logger.Error("Error trying to send Nack", zap.Error(err))
	// 	}
	// } else {
	// 	logger.Info("Sending Acknowledge")
	// 	logger.Info("Delivery tag ack", zap.Uint64("delivery_tag", delivery.DeliveryTag))
	// 	err = videoCh.Ack(delivery.DeliveryTag, false)
	// 	if err != nil {
	// 		logger.Error("Error trying to send Ack", zap.Error(err))
	// 	}
	// }

	if err != nil {
		// NOTE: If something fails I create the message again manually (I still don't know why the channel closes after send ack)
		// NOTE: This is a poor solution
		logger.Error("Error - preparing Nack", zap.Error(err))
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
			"",                  // exchange
			cfgQueue.VideoQueue, // routing key
			false,               // mandatory
			false,               // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         body,
			})

		if err != nil {
			logger.Error("Error sending the nack message to the queue", zap.Error(err))
		}
	}

	return
}

func getMp3Buffer(buf []byte) ([]byte, error) {
	logger := configLogger.GetLogger()
	// TODO: File buffer
	// buf, err := os.ReadFile("./Dortmund.mp4")

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer

	command := exec.Command("ffmpeg", "-i", "pipe:0", "-b:a", "320K", "-vn", "-f", "mp3", "pipe:1")

	command.Stdout = bufio.NewWriter(&stdOut)
	command.Stderr = bufio.NewWriter(&stdErr)

	stdin, err := command.StdinPipe()
	if err != nil {
		logger.Error("Error creating stdin pipe", zap.Error(err))
		return nil, err
	}

	err = command.Start()
	if err != nil {
		logger.Error("Error trying to start mp3 conversion command", zap.Error(err))
		return nil, err
	}

	_, err = stdin.Write(buf)
	if err != nil {
		logger.Error("Error trying to write to stdin", zap.Error(err))
		return nil, err
	}

	err = stdin.Close()
	if err != nil {
		logger.Error("Error trying to close stdin pipe", zap.Error(err))
		return nil, err
	}

	err = command.Wait()
	if err != nil {
		logger.Error("Error waiting for mp3 conversion command to complete", zap.Error(err))
		return nil, err
	}

	// fmt.Println("LOG", string(stdErr.Bytes()))
	// fmt.Println("OUTPUT", stdOut.Bytes())
	return stdOut.Bytes(), nil
}

func start(queueMessage QueueMessage, fsVideos *gridfs.Bucket, fsMp3 *gridfs.Bucket, ch *amqp.Channel) error {
	cfgQueue := configQueue.GetQueueConfig()
	logger := configLogger.GetLogger()

	fileBuffer := bytes.NewBuffer(nil)

	id, err := primitive.ObjectIDFromHex(queueMessage.VideoFId)
	if err != nil {
		logger.Error("Error trying to parse the hex to object id instance", zap.Error(err))
		return err
	}

	if _, err := fsVideos.DownloadToStream(id, fileBuffer); err != nil {
		logger.Error("Error opening the stream to the videos database", zap.Error(err))
		return err
	}

	audio, err := getMp3Buffer(fileBuffer.Bytes())
	if err != nil {
		logger.Error("Error trying to get the mp3 from video file", zap.Error(err))
		return err
	}

	fileName := fmt.Sprintf("%s.mp3", queueMessage.VideoFId)

	fileId, err := fsMp3.UploadFromStream(fileName, bytes.NewReader(audio))

	hexId := fileId.String()[10:34]

	newMessage := QueueMessage{
		VideoFId:  queueMessage.VideoFId,
		Mp3FileId: hexId,
		Username:  queueMessage.Username,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(newMessage)
	if err != nil {
		logger.Error("Error parsing the json queue message body", zap.Error(err))
		return err
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
		logger.Error("Error sending the message to the queue", zap.Error(err))
		if err := fsMp3.Delete(fileId); err != nil {
			logger.Error("Error trying to delete file in mongoDB", zap.Error(err))
			return err
		}
	}

	logger.Info("mp3 successfully loaded in mongoDB database", zap.String("mp3_file_name", fileName))
	return nil
}
