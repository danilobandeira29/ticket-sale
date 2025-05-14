package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"strconv"
)

type EventSectionID = domain.UUID

type EventSection struct {
	entity             domain.Entity
	ID                 EventSectionID
	Name               string
	Description        *string
	IsPublished        bool
	TotalSpots         int32
	TotalSpotsReserved int32
	Price              float64
	Spots              EventSpotSet
}

type EventSpotSet = domain.Set[string, *EventSpot]

type EventSectionProps struct {
	ID                             *EventSectionID
	Name                           string
	Description                    *string
	TotalSpots, TotalSpotsReserved int32
	IsPublished                    bool
	Price                          float64
	Spots                          EventSpotSet
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
		Description:        props.Description,
		IsPublished:        props.IsPublished,
		TotalSpots:         props.TotalSpots,
		TotalSpotsReserved: props.TotalSpotsReserved,
		Price:              props.Price,
		Spots:              *domain.NewSet[string, *EventSpot](),
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
	section := &EventSection{
		entity:             domain.Entity{},
		ID:                 *id,
		Name:               command.Name,
		Description:        command.Description,
		IsPublished:        false,
		TotalSpots:         command.TotalSpots,
		TotalSpotsReserved: 0,
		Price:              command.Price,
		Spots:              *domain.NewSet[string, *EventSpot](),
	}
	for i := range section.TotalSpots {
		idx := strconv.Itoa(int(i + 1))
		spot, errEventSpot := CreateEventSpot(idx)
		if errEventSpot != nil {
			return nil, fmt.Errorf("create event section: creating spots: %v", errEventSpot)
		}
		section.Spots.Add(idx, spot)
	}
	return section, nil
}

func (e *EventSection) String() string {
	return e.entity.String(e)
}

func (e *EventSection) Publish() {
	e.IsPublished = true
}

func (e *EventSection) Unpublish() {
	e.IsPublished = false
}

func (e *EventSection) PublishAll() {
	e.Publish()
	for _, s := range e.Spots.Data {
		s.Publish()
	}
}

func (e *EventSection) UnpublishAll() {
	e.Unpublish()
	for _, s := range e.Spots.Data {
		s.Unpublish()
	}
}
