package repository

import (
	"context"
	"encoding/json"
	"event-sourcing/internal/database/order/queries"
	"event-sourcing/internal/domain"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	OrderRepository struct {
		db *sqlx.DB
	}

	Event struct {
		ID        uuid.UUID       `db:"aggregate_id"`
		Version   uint            `db:"version"`
		EventType string          `db:"event_type"`
		EventData json.RawMessage `db:"event_data"`
		CreatedAt time.Time       `db:"created_at"`
	}
)

func FromDomain(order domain.Order) []Event {
	res := []Event{}
	version := order.Version
	for _, e := range order.Events {
		version++
		res = append(res,
			Event{
				ID:        order.ID,
				Version:   version,
				EventType: e.GetName(),
				EventData: e.GetBytes(),
				CreatedAt: time.Now(),
			},
		)
	}

	return res
}

func (e *Event) ToDomain() domain.Event {
	var res domain.Event
	switch e.EventType {
	case "ORDER_PLACED":
		var data domain.OrderPlacedData
		_ = json.Unmarshal(e.EventData, &data)
		res = domain.OrderPlaced{
			ID:   e.ID,
			Name: e.EventType,
			Data: data,
		}
	case "PAYMENT_CAPTURED":
		var data domain.PaymentCaptuedData
		_ = json.Unmarshal(e.EventData, &data)
		res = domain.PaymentCaptured{
			ID:   e.ID,
			Name: e.EventType,
			Data: data,
		}
	}

	return res
}

func NewRepository(db *sqlx.DB) OrderRepository {
	return OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Find(ctx context.Context, orderID uuid.UUID) (domain.Order, error) {
	order := domain.NewOrder(orderID)
	rows, err := r.db.Unsafe().QueryxContext(ctx, queries.ListEvent, orderID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return order, err
	}

	events := []domain.Event{}
	for rows.Next() {
		var event Event
		err := rows.StructScan(&event)
		if err != nil {
			fmt.Println("error scaning row", err)
			return order, err
		}

		events = append(events, event.ToDomain())
	}

	for _, e := range events {
		order.When(e)
		order.IncrementVersion()
	}

	return order, nil
}

func (r *OrderRepository) Save(ctx context.Context, order domain.Order) error {
	eventsDTO := FromDomain(order)
	for _, e := range eventsDTO {
		_, err := r.save(ctx, e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *OrderRepository) save(ctx context.Context, event Event) (Event, error) {
	q, args, err := sqlx.Named(queries.Save, event)
	if err != nil {
		return event, err
	}
	q = r.db.Rebind(q)

	err = r.db.Unsafe().QueryRowxContext(ctx, q, args...).StructScan(&event)
	if err != nil {
		return event, err
	}

	return event, nil
}
