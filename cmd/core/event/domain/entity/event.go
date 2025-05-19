package entity

import (
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type EventID = domain.UUID

type Event struct {
	aggregate          domain.AggregateRoot
	ID                 EventID
	Name               string
	Description        *string
	Date               time.Time
	IsPublished        bool
	TotalSpots         int32
	TotalSpotsReserved int32
	PartnerID          PartnerID
	Sections           EventSectionSet
}

type EventSectionSet = domain.Set[string, *EventSection]

type EventProps struct {
	ID                             *EventID
	Name                           string
	Description                    *string
	TotalSpots, TotalSpotsReserved int32
	PartnerID                      *PartnerID
	IsPublished                    bool
	Date                           time.Time
	EventSectionSet                EventSectionSet
}

func NewEvent(props EventProps) (*Event, error) {
	var id EventID
	if props.ID != nil {
		id = *props.ID
	} else {
		i, err := domain.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("new event: creating id: %v", err)
		}
		id = *i
	}
	var partnerID PartnerID
	if props.PartnerID != nil {
		partnerID = *props.PartnerID
	} else {
		i, err := domain.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("new event: creating partner id: %v", err)
		}
		partnerID = *i
	}
	return &Event{
		aggregate:          domain.AggregateRoot{},
		ID:                 id,
		Name:               props.Name,
		Description:        props.Description,
		Date:               props.Date,
		IsPublished:        props.IsPublished,
		TotalSpots:         props.TotalSpots,
		TotalSpotsReserved: props.TotalSpotsReserved,
		PartnerID:          partnerID,
		Sections:           props.EventSectionSet,
	}, nil
}

type CreateEventCommand struct {
	Name        string
	Description *string
	Date        time.Time
	PartnerID   PartnerID
}

func CreateEvent(command CreateEventCommand) (*Event, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("create event: creating id: %v", err)
	}
	return &Event{
		aggregate:          domain.AggregateRoot{},
		ID:                 *id,
		Name:               command.Name,
		Description:        command.Description,
		Date:               command.Date,
		IsPublished:        false,
		TotalSpots:         0,
		TotalSpotsReserved: 0,
		PartnerID:          command.PartnerID,
		Sections:           *domain.NewSet[string, *EventSection](),
	}, nil
}

type AddSectionCommand struct {
	Name        string
	Description *string
	TotalSpots  int32
	Price       float64
}

func (e *Event) AddSection(command AddSectionCommand) error {
	section, err := CreateEventSection(CreateEventSectionCommand{
		Name:        command.Name,
		Description: command.Description,
		TotalSpots:  command.TotalSpots,
		Price:       command.Price,
	})
	if err != nil {
		return fmt.Errorf("add section: %v", err)
	}
	e.Sections.Add(section.ID.String(), section)
	e.TotalSpots += section.TotalSpots
	return nil
}

func (e *Event) String() string {
	return e.aggregate.String(e)
}

func (e *Event) ChangeDate(now, t time.Time) error {
	if now.After(t) {
		return fmt.Errorf("it is not possible to change the event date to a past date")
	}
	e.Date = t
	return nil
}

func (e *Event) Publish() {
	e.IsPublished = true
}

func (e *Event) Unpublish() {
	e.IsPublished = false
}

func (e *Event) PublishAll() {
	e.Publish()
	for _, s := range e.Sections.Data {
		s.PublishAll()
	}
}

func (e *Event) UnpublishAll() {
	e.Unpublish()
	for _, s := range e.Sections.Data {
		s.UnpublishAll()
	}
}

type AllowReserveSpotInput struct {
	SectionID EventSectionID
	SpotID    EventSpotID
}

func (e *Event) AllowReserveSpot(input AllowReserveSpotInput) (bool, error) {
	if !e.IsPublished {
		return false, errors.New("event allow reserve spot: event not published")
	}
	var section *EventSection
	for _, s := range e.Sections.Data {
		if s.ID.Equal(input.SectionID) {
			section = s
			break
		}
	}
	if section == nil {
		return false, errors.New("event allow reserve spot: section not found")
	}
	return section.AllowReserveSpot(input.SpotID)
}

func (e *Event) Section(id EventSectionID) (EventSection, error) {
	for _, s := range e.Sections.Data {
		if s.ID.Equal(id) {
			return *s, nil
		}
	}
	return EventSection{}, errors.New("section not found")
}

type ChangeSectionInput struct {
	SectionID         EventSectionID
	Name, Description *string
}

func (e *Event) ChangeSectionInfo(command ChangeSectionInput) error {
	if exists := e.Sections.Exists(command.SectionID.String()); !exists {
		return errors.New("event change section info: section not found")
	}
	for _, s := range e.Sections.Data {
		if s.ID != command.SectionID {
			continue
		}
		if command.Name != nil {
			s.ChangeName(*command.Name)
		}
		if command.Description != nil {
			s.ChangeDescription(*command.Description)
		}
	}
	return nil
}
