package application

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/infra/db"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/infra/unitofwork"
	"testing"
	"time"
)

func TestEventService_Create(t *testing.T) {
	conn, _ := db.PostgresConn()
	partnerRepo := db.NewPartnerRepository(conn)
	eventRepo := db.NewEventRepository(conn)
	uow := unitofwork.NewUoW(conn)
	uow.RegisterFactory("PartnerRepository", func(exec db.Executor) any {
		return db.NewPartnerRepository(exec)
	})
	uow.RegisterFactory("EventRepository", func(exec db.Executor) any {
		return db.NewEventRepository(exec)
	})
	service := NewEventService(eventRepo, partnerRepo, uow)
	partner, _ := entity.CreatePartner("Partner 1")
	_ = partnerRepo.Save(partner)
	date, _ := time.Parse(time.RFC3339, "2030-01-01T10:00:00-03:00")
	eventName := time.Now().Format(time.RFC3339)
	err := service.Create(CreateInput{
		Name:        eventName,
		Date:        date,
		PartnerID:   partner.ID.String(),
		Description: nil,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	result, _ := eventRepo.FindAll()
	var event *entity.Event
	for _, r := range result {
		if r.Name == eventName {
			event = r
			break
		}
	}
	if event == nil {
		t.Error("expected event to be created")
		return
	}
}
