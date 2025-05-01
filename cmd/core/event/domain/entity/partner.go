package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type PartnerID = domain.UUID

type Partner struct {
	ID PartnerID
}

func NewPartner(n string) (*Partner, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("new partner: %v", err)
	}
	return &Partner{ID: *id}, nil
}
