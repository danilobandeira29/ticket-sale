package db

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"log"
	"testing"
	"time"
)

func TestEventRepository_Save(t *testing.T) {
	conn, _ := PostgresConn()
	tx, _ := conn.Begin()
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback transaction: %v", err)
		}
	}()
	repo := NewEventRepository(tx)
	eventTime, _ := time.Parse(time.RFC3339, "2026-01-01T10:00:00-03:00")
	partnerRepo := NewPartnerRepository(tx)
	partner, _ := entity.CreatePartner("Partner")
	_ = partnerRepo.Save(partner)
	event, _ := partner.CreateEvent(entity.PartnerCreateEvent{
		Name:        "Event 1",
		Description: nil,
		Date:        eventTime,
	})
	_ = event.AddSection(entity.AddSectionCommand{
		Name:        "Premium",
		Description: nil,
		TotalSpots:  1000,
		Price:       999.00,
	})
	err := repo.Save(event)
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
}
