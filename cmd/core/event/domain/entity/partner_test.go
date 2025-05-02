package entity

import (
	"testing"
	"time"
)

func TestPartner_CreateEvent(t *testing.T) {
	p, err := CreatePartner("Parceiro 1")
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
	_, err = p.CreateEvent(PartnerCreateEvent{
		Name:        "New Event",
		Description: nil,
		Date:        time.Now(),
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v\n", err)
		return
	}
}
