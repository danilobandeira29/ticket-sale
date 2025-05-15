package application

import (
	"fmt"
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
	eventDescription := "Nice description"
	eventID, err := service.Create(CreateInput{
		Name:        "Event name " + eventName,
		Date:        date,
		PartnerID:   partner.ID.String(),
		Description: &eventDescription,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	result, err := eventRepo.FindByID(*eventID)
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	if result.ID.String() != *eventID {
		t.Errorf("event not created\nexpected: %s\ngot: %s", *eventID, result.ID.String())
		return
	}
	fmt.Printf("%#v", result)
}

//func TestEventService_AddSection(t *testing.T) {
//	conn, _ := db.PostgresConn()
//	partnerRepo := db.NewPartnerRepository(conn)
//	eventRepo := db.NewEventRepository(conn)
//	uow := unitofwork.NewUoW(conn)
//	uow.RegisterFactory("PartnerRepository", func(exec db.Executor) any {
//		return db.NewPartnerRepository(exec)
//	})
//	uow.RegisterFactory("EventRepository", func(exec db.Executor) any {
//		return db.NewEventRepository(exec)
//	})
//	service := NewEventService(eventRepo, partnerRepo, uow)
//	partner, _ := entity.CreatePartner("Partner 22")
//	_ = partnerRepo.Save(partner)
//	date, _ := time.Parse(time.RFC3339, "2030-01-01T10:00:00-03:00")
//	eventName := time.Now().Format(time.RFC3339)
//	eventID, err := service.Create(CreateInput{
//		Name:        "Event name " + eventName,
//		Date:        date,
//		PartnerID:   partner.ID.String(),
//		Description: nil,
//	})
//	if err != nil {
//		t.Errorf("expected error to be empty\ngot: %v", err)
//		return
//	}
//}
