package event

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type PartnerCreated struct {
	domain.DEvent
}

func NewPartnerCreated(id string, occurredAt time.Time, version uint) *PartnerCreated {
	e := domain.NewEvent(id, "partner created", occurredAt, version)
	return &PartnerCreated{
		*e,
	}
}
