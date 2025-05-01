package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type EventSectionID = domain.UUID

type EventSection struct {
	entity             domain.Entity
	ID                 EventSectionID
	Name, Description  string
	IsPublished        bool
	TotalSpots         int32
	TotalSpotsReserved int32
	Price              float64
}

type EventSectionProps struct {
	ID                             *EventSectionID
	Name                           string
	Description                    *string
	TotalSpots, TotalSpotsReserved int32
	IsPublished                    bool
}

func NewEventSection(props EventSectionProps) (*EventSection, error) {
	var id EventSectionID
	if props.ID != nil {
		id = *props.ID
	} else {
		i, err := domain.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("new event section: creating id: %v", err)
		}
		id = *i
	}
	return &EventSection{
		entity:             domain.Entity{},
		ID:                 id,
		Name:               props.Name,
		Description:        *props.Description,
		IsPublished:        props.IsPublished,
		TotalSpots:         props.TotalSpots,
		TotalSpotsReserved: props.TotalSpotsReserved,
	}, nil
}

type CreateEventSectionCommand struct {
	Name        string
	Description *string
	TotalSpots  int32
	Price       float64
}

func CreateEventSection(command CreateEventSectionCommand) (*EventSection, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("create event section: creating id: %v", err)
	}
	return &EventSection{
		entity:             domain.Entity{},
		ID:                 *id,
		Name:               command.Name,
		Description:        *command.Description,
		IsPublished:        false,
		TotalSpots:         command.TotalSpots,
		TotalSpotsReserved: 0,
		Price:              command.Price,
	}, nil
}
