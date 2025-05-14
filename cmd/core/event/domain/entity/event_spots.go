package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type EventSpotID = domain.UUID

type EventSpot struct {
	entity      domain.Entity
	ID          EventSpotID
	Location    *string
	IsPublished bool
	IsReserved  bool
}

type EventSpotProps struct {
	ID          *EventSpotID
	Location    *string
	IsReserved  bool
	IsPublished bool
}

func NewEventSpot(props EventSpotProps) (*EventSpot, error) {
	var id EventSpotID
	if props.ID != nil {
		id = *props.ID
	} else {
		i, err := domain.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("new event spots: creating id: %v", err)
		}
		id = *i
	}
	return &EventSpot{
		entity:      domain.Entity{},
		ID:          id,
		Location:    props.Location,
		IsPublished: props.IsPublished,
		IsReserved:  props.IsReserved,
	}, nil
}

func CreateEventSpot(loc string) (*EventSpot, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("create event spot: creating id: %v", err)
	}
	return &EventSpot{
		entity:      domain.Entity{},
		ID:          *id,
		Location:    &loc,
		IsPublished: false,
		IsReserved:  false,
	}, nil
}

func (e *EventSpot) String() string {
	return e.entity.String(e)
}

func (e *EventSpot) Publish() {
	e.IsPublished = true
}

func (e *EventSpot) Unpublish() {
	e.IsPublished = false
}
