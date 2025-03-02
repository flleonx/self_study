package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type EventName string

const (
	OrderPlacedName     EventName = "ORDER_PLACED"
	PaymentCapturedName EventName = "PAYMENT_CAPTURED"
)

type (
	OrderPlacedData struct {
		ShippingAddress string
		ShippingMethod  string
	}

	OrderPlaced struct {
		ID   uuid.UUID
		Name string
		Data OrderPlacedData
	}

	PaymentCaptuedData struct {
		PaidAt time.Time
	}

	PaymentCaptured struct {
		ID   uuid.UUID
		Name string
		Data PaymentCaptuedData
	}
)

func NewOrderPlaced(id uuid.UUID, data OrderPlacedData) OrderPlaced {
	return OrderPlaced{
		ID:   id,
		Name: "ORDER_PLACED",
		Data: data,
	}
}

func NewPaymentCaptured(id uuid.UUID, data PaymentCaptuedData) PaymentCaptured {
	return PaymentCaptured{
		ID:   id,
		Name: "PAYMENT_CAPTURED",
		Data: data,
	}
}

func (o OrderPlaced) GetID() uuid.UUID {
	return o.ID
}

func (o OrderPlaced) GetName() string {
	return o.Name
}

func (o OrderPlaced) GetBytes() []byte {
	bytes, _ := json.Marshal(o.Data)

	return bytes
}

func (o PaymentCaptured) GetID() uuid.UUID {
	return o.ID
}

func (o PaymentCaptured) GetName() string {
	return o.Name
}

func (o PaymentCaptured) GetBytes() []byte {
	bytes, _ := json.Marshal(o.Data)

	return bytes
}

type Event interface {
	GetID() uuid.UUID
	GetName() string
	GetBytes() []byte
}
