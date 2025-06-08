package application

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type PartnerService struct {
	repo    domain.Repository[entity.Partner]
	service *application.Service
	uow     application.UnitOfWork
}

func NewPartnerService(repo domain.Repository[entity.Partner], uow application.UnitOfWork, manager domain.EventManager) *PartnerService {
	return &PartnerService{repo: repo, service: application.NewService(uow, manager)}
}

func (p *PartnerService) Create(name string) (*entity.Partner, error) {
	partner, err := entity.CreatePartner(name, time.Now())
	if err != nil {
		return nil, fmt.Errorf("service create: %w", err)
	}
	_, err = p.service.Run([]*domain.AggregateRoot{partner.Aggregate}, func() any {
		return p.repo.Save(partner)
	})
	if err != nil {
		return nil, err
	}
	return partner, nil
}
