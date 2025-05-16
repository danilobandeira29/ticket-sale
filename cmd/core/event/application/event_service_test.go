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
	eventDescription := "Nice description"
	eventID, err := service.Create(CreateInput{
		Name:        "Create" + time.Now().Format(time.RFC3339),
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
}

func TestEventService_AddSection(t *testing.T) {
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
	event, _ := partner.CreateEvent(entity.PartnerCreateEvent{
		Name:        "AddSection " + time.Now().Format(time.RFC3339),
		Description: nil,
		Date:        time.Now(),
	})
	_ = eventRepo.Save(event)
	spotDescription := "Premium assets"
	err := service.AddSection(AddSectionInput{
		Name:        "Premium",
		Description: &spotDescription,
		TotalSpots:  100,
		Price:       999.98,
		EventID:     event.ID.String(),
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	e, err := eventRepo.FindByID(event.ID.String())
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	if e.TotalSpots != 100 {
		t.Errorf("wrong total spot\nexpected: 100\ngot: %d", e.TotalSpots)
		return
	}
}

func TestEventService_ChangeSectionInfo(t *testing.T) {
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
	event, _ := partner.CreateEvent(entity.PartnerCreateEvent{
		Name:        "AddSection " + time.Now().Format(time.RFC3339),
		Description: nil,
		Date:        time.Now(),
	})
	_ = event.AddSection(entity.AddSectionCommand{
		Name:        "Premium",
		Description: nil,
		TotalSpots:  100,
		Price:       1999.97,
	})
	var sectionID string
	for id := range event.Sections.Data {
		sectionID = id
		break
	}
	_ = eventRepo.Save(event)
	sectionDesc := "New Desc"
	sectionName := "New Name"
	err := service.ChangeSectionInfo(ChangeSectionInfo{
		EventID:     event.ID.String(),
		SectionID:   sectionID,
		Name:        &sectionName,
		Description: &sectionDesc,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	e, err := eventRepo.FindByID(event.ID.String())
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	if !e.Sections.Exists(sectionID) {
		t.Error("section not founded")
		return
	}
	if e.Sections.Data[sectionID].Name != sectionName {
		t.Errorf("wrong name\nexpected: %s\ngot: %s", sectionName, e.Sections.Data[sectionID].Name)
		return
	}
	if *e.Sections.Data[sectionID].Description != sectionDesc {
		t.Errorf("wrong description \nexpected: %s\ngot: %s", sectionDesc, *e.Sections.Data[sectionID].Description)
		return
	}
}
