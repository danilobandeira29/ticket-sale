package entity

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type SpotReservation struct {
	aggregate       domain.AggregateRoot
	SpotID          EventSpotID
	ReservationDate time.Time
	CustomerID      CustomerID
}

func NewSpotReservation(spotID EventSpotID, date time.Time, customerID CustomerID) *SpotReservation {
	return &SpotReservation{
		SpotID:          spotID,
		ReservationDate: date,
		CustomerID:      customerID,
	}
}

func (s *SpotReservation) String() string {
	return s.aggregate.String(s)
}

type SpotReservationCreateCommand struct {
	SpotID          EventSpotID
	CustomerID      CustomerID
	ReservationDate time.Time
}

func CreateSpotReservation(command SpotReservationCreateCommand) *SpotReservation {
	return &SpotReservation{
		aggregate:       domain.AggregateRoot{},
		SpotID:          command.SpotID,
		ReservationDate: command.ReservationDate,
		CustomerID:      command.CustomerID,
	}
}

func (s *SpotReservation) ChangeReservation(cID CustomerID, date time.Time) {
	s.CustomerID = cID
	s.ReservationDate = date
}
