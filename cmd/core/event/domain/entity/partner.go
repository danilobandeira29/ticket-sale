package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type PartnerID = domain.UUID

type Partner struct {
	Name string
	ID   PartnerID
}

func CreatePartner(n string) (*Partner, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("new partner: %v", err)
	}
	return &Partner{ID: *id, Name: n}, nil
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
