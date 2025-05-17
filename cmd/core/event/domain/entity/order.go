package entity

import "github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"

type OrderID = domain.UUID
type Order struct {
	aggregate   domain.AggregateRoot
	ID          OrderID
	CustomerID  CustomerID
	Amount      float64
	EventSpotID EventSpotID
	Status      OrderStatus
}

type OrderProps struct {
	ID          OrderID
	CustomerID  CustomerID
	Amount      float64
	EventSpotID EventSpotID
}

type OrderStatus int

const (
	StatusPending OrderStatus = iota
)

var orderStatus = map[OrderStatus]string{
	StatusPending: "pending",
}

func (o OrderStatus) String() string {
	return orderStatus[o]
}

func NewOrder(props OrderProps) *Order {
	return &Order{
		aggregate:   domain.AggregateRoot{},
		ID:          props.ID,
		CustomerID:  props.CustomerID,
		Amount:      props.Amount,
		EventSpotID: props.EventSpotID,
		Status:      StatusPending,
	}
}

func (o *Order) String() string {
	return o.aggregate.String(o)
}
