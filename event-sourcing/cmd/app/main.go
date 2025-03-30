package main

import (
	"context"
	"event-sourcing/configuration"
	"event-sourcing/internal/database"
	"event-sourcing/internal/domain"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"

	orders "event-sourcing/internal/database/order"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	conn, err := amqp.Dial("amqp://guest:guest@broker:5672/")
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err, "failed to open a channel")
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err, "error creating queue")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := configuration.LoadConfig()
	if err != nil {
		fmt.Println("error loading config", err)
	}

	db, err := database.NewDatabase(config)
	if err != nil {
		fmt.Println("error loading config", err)
	}

	database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	orderRepo := orders.NewRepository(db)

	order := domain.Order{
		ID:              uuid.New(),
		ShippingAddress: "EMPTY",
		ShippingMethod:  "EMPTY",
		PlacedAt:        time.Now(),
		ShippedAt:       time.Now(),
		Status:          "PENDING",
		PaidAt:          time.Now(),
		Events: []domain.Event{
			domain.NewOrderPlaced(
				uuid.New(),
				domain.OrderPlacedData{
					ShippingAddress: "Test 1",
					ShippingMethod:  "Bus",
				},
			),
			domain.NewPaymentCaptured(
				uuid.New(),
				domain.PaymentCaptuedData{
					PaidAt: time.Now(),
				},
			),
		},
	}

	err = orderRepo.Save(context.Background(), order)
	if err != nil {
		slog.Error("ORDER SAVE", slog.Any("error", err))
		os.Exit(1)
	}

	order, err = orderRepo.Find(context.Background(), order.ID)
	if err != nil {
		slog.Error("ORDER FIND", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("LAST ORDER", slog.Any("order", order))

	body := "Hello World!"
	for {
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		slog.Info(fmt.Sprintf(" [x] Sent %s\n", body))
		time.Sleep(20 * time.Second)
	}
}
