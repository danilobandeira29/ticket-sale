package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/event"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type PartnerID = domain.UUID

type Partner struct {
	Aggregate *domain.AggregateRoot
	Name      string
	ID        PartnerID
}

func CreatePartner(n string, now time.Time) (*Partner, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("new partner: %v", err)
	}
	p := &Partner{ID: *id, Name: n, Aggregate: domain.NewAggregateRoot()}
	p.Aggregate.AddEvent(event.NewPartnerCreated(id.String(), now, 1))
	return p, nil
}

type PartnerCreateEvent struct {
	Name        string
	Description *string
	Date        time.Time
}

func (p *Partner) CreateEvent(command PartnerCreateEvent) (*Event, error) {
	return CreateEvent(CreateEventCommand{
		Name:        command.Name,
		Description: command.Description,
		Date:        command.Date,
		PartnerID:   p.ID,
	})
}
