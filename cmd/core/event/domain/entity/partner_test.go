package entity

import (
	"testing"
	"time"
)

func TestPartner_CreateEvent(t *testing.T) {
	p, _ := CreatePartner("Parceiro 1", time.Now())
	_, err := p.CreateEvent(PartnerCreateEvent{
		Name:        "New Event",
		Description: nil,
		Date:        time.Now(),
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
}

func TestCreatePartner(t *testing.T) {
	p, err := CreatePartner("Partner 2", time.Now())
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
	if p == nil {
		t.Errorf("expected partner to be created\ngot: nil")
		return
	}
}
