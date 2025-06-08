package application

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type Service struct {
	uow          UnitOfWork
	eventManager domain.EventManager
	err          error
}

func NewService(uow UnitOfWork, manager domain.EventManager) *Service {
	return &Service{
		uow:          uow,
		eventManager: manager,
	}
}

func (s *Service) Finish(aggregates []*domain.AggregateRoot) error {
	if err := s.uow.Commit(); err != nil {
		return fmt.Errorf("finish: %w", err)
	}
	for _, agg := range aggregates {
		s.eventManager.Publish(agg)
	}
	return nil
}

func (s *Service) Run(aggregates []*domain.AggregateRoot, fn func() any) (any, error) {
	s.Start()
	result := fn()
	if err := s.Finish(aggregates); err != nil {
		return result, err
	}
	return result, nil
}

func (s *Service) Start() {

}
