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
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"

	orders "event-sourcing/internal/database/order"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

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

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
