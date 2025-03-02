package domain

import (
	"time"

	"github.com/google/uuid"
)

type (
	Order struct {
		ID              uuid.UUID
		ShippingAddress string
		ShippingMethod  string
		PlacedAt        time.Time
		ShippedAt       time.Time
		Status          string
		PaidAt          time.Time
		Events          []Event
		Version         uint
	}
)

func NewOrder(id uuid.UUID) Order {
	return Order{
		ID: id,
	}
}

func (o *Order) When(event Event) {
	switch v := event.(type) {
	case OrderPlaced:
		o.whenOrderPlaced(v)
	case PaymentCaptured:
		o.whenPaymentCaptured(v)
	}
}

func (o *Order) IncrementVersion() {
	o.Version++
}

func (o *Order) whenPaymentCaptured(event PaymentCaptured) {
	o.PaidAt = event.Data.PaidAt
	o.updateStatus(event)
}

func (o *Order) whenOrderPlaced(event OrderPlaced) {
	o.ShippingAddress = event.Data.ShippingAddress
	o.ShippingMethod = event.Data.ShippingMethod
	o.updateStatus(event)
}

func (o *Order) updateStatus(event any) {
	switch event.(type) {
	case OrderPlaced:
		o.Status = "PLACED"
	case PaymentCaptured:
		o.Status = "PAID"
	}
}
