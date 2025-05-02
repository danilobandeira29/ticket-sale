package entity

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"testing"
	"time"
)

func TestCreateEvent_WithSectionAndSpot(t *testing.T) {
	partnerID, _ := domain.NewUUID()
	event, err := CreateEvent(CreateEventCommand{
		Name:        "Golang is easy",
		Description: nil,
		Date:        time.Now(),
		PartnerID:   *partnerID,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	if len(event.Sections.Data) != 0 {
		t.Errorf("expected event's section to have len 0")
		return
	}
	section, err := CreateEventSection(CreateEventSectionCommand{
		Name:        "Section name",
		Description: nil,
		TotalSpots:  0,
		Price:       100.44,
	})
	if err != nil {
		t.Errorf("expected error create event section to be empty\ngot: %v", err)
		return
	}
	if len(section.Spots.Data) != 0 {
		t.Errorf("expected section's spots to have len 0")
		return
	}
	spot, err := CreateEventSpot()
	if err != nil {
		t.Errorf("expected error create event spot to be empty\bgot: %v", err)
		return
	}
	section.Spots.Add(spot.ID.String(), *spot)
	if !section.Spots.Exists(spot.ID.String()) {
		t.Errorf("expected spot to exists in section")
		return
	}
	event.Sections.Add(section.ID.String(), *section)
	if !event.Sections.Exists(section.ID.String()) {
		t.Errorf("expected section to exists in event")
		return
	}
}

func TestEvent_AddSection(t *testing.T) {
	partnerID, _ := domain.NewUUID()
	event, err := CreateEvent(CreateEventCommand{
		Name:        "Event Add Section",
		Description: nil,
		Date:        time.Now(),
		PartnerID:   *partnerID,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	if event.Sections.Size() != 0 {
		t.Errorf("section size expected: 0\ngot: %d\n", event.Sections.Size())
		return
	}
	err = event.AddSection(AddSectionCommand{
		Name:        "",
		Description: nil,
		TotalSpots:  10,
		Price:       66,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
	if event.TotalSpots != 10 {
		t.Errorf("section size expected: 10\ngot: %d\n", event.TotalSpots)
		return
	}
	err = event.AddSection(AddSectionCommand{
		Name:        "Premium",
		Description: nil,
		TotalSpots:  5,
		Price:       99,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
	if event.TotalSpots != 15 {
		t.Errorf("section size expected: 15\ngot: %d\n", event.TotalSpots)
		return
	}
}

func TestCreateEvent_SpotsSize(t *testing.T) {
	partnerID, _ := domain.NewUUID()
	event, err := CreateEvent(CreateEventCommand{
		Name:        "Event Spots Size",
		Description: nil,
		Date:        time.Now(),
		PartnerID:   *partnerID,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	if event.Sections.Size() != 0 {
		t.Errorf("section size expected: 0\ngot: %d\n", event.Sections.Size())
		return
	}
	err = event.AddSection(AddSectionCommand{
		Name:        "Basic",
		Description: nil,
		TotalSpots:  10,
		Price:       66,
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
	for _, v := range event.Sections.Data {
		if v.Spots.Size() != 10 {
			t.Errorf("section spots size: expected: 10\ngot: %d\n", v.Spots.Size())
			return
		}
	}
}
