package db

import "github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"

type SpotReservationRepository struct {
	executor Executor
}

func NewSpotReservationRepository(executor Executor) *SpotReservationRepository {
	return &SpotReservationRepository{executor: executor}
}

// todo create table spot_reservations
// todo create table order

func (s *SpotReservationRepository) Save(t *entity.SpotReservation) error {
	//TODO implement me
	panic("implement me")
}

func (s *SpotReservationRepository) FindByID(id any) (*entity.SpotReservation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpotReservationRepository) FindAll() ([]*entity.SpotReservation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SpotReservationRepository) Delete(id any) error {
	//TODO implement me
	panic("implement me")
}
