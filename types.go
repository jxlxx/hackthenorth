package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type OrderType string

const (
	BuyLimit   OrderType = "buy_limit"
	SellLimit  OrderType = "sell_limit"
	BuyMarket  OrderType = "buy_market"
	SellMarket OrderType = "sell_market"
	AllTypes   OrderType = "*"
)

type OrderStatus string

const (
	Open      OrderStatus = "open"
	Closed    OrderStatus = "closed"
	Cancelled OrderStatus = "cancelled"
	Failed    OrderStatus = "failed"
	AllStatus OrderStatus = "*"
)

type Order struct {
	TimeStamp time.Time
	ID        uuid.UUID
	TraderID  uuid.UUID
	Exchange  string
	Security  string
	OrderType OrderType
	Status    OrderStatus
	Quantity  int
	Price     Cents // In Cents
}

type Cents int
type Dollars int

func (c Cents) Bytes() []byte {
	b, _ := json.Marshal(c)
	return b
}

func (c Cents) Dollars() (Dollars, Cents) {
	leftovers := int(c) % 100
	return Dollars(int(c) / 100), Cents(leftovers)
}

type Bot struct {
	name string
	ID   uuid.UUID
	nc   *nats.Conn
}

type Trader struct {
	ID       uuid.UUID
	Name     string
	Password string
}

type NewOrder struct {
	ExchangeID string
	OrderType  OrderType
	SecurityID string
	Qty        int
}

type Transaction struct {
	OrderType  string
	TraderID   uuid.UUID
	Security   string
	Quantity   int
	Price      Cents
	TotalPrice Cents
	Networth   Cents
	Time       time.Time
}

type Portfolio struct {
	Funds           Cents
	SecurityHolding map[string]int
}

func (o NewOrder) String() string {
	exchange := o.ExchangeID
	if exchange == "" {
		exchange = "*"
	}
	security := o.SecurityID
	if security == "" {
		security = "*"
	}
	orderType := o.OrderType
	if orderType == "" {
		orderType = "*"
	}
	qty := fmt.Sprintf("%d", o.Qty)
	return fmt.Sprintf("orders.new.%s.%s.%s.%s", exchange, orderType, security, qty)
}

type OrderPayload struct {
	TraderID uuid.UUID
	Price    Cents
}

func (o OrderPayload) Bytes() []byte {
	b, _ := json.Marshal(o)
	return b
}
