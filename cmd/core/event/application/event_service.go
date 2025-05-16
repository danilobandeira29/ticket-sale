package application

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"time"
)

type EventService struct {
	eventRepo   domain.Repository[entity.Event]
	partnerRepo domain.Repository[entity.Partner]
	uow         application.UnitOfWork
}

func NewEventService(eventRepo domain.Repository[entity.Event], partnerRepo domain.Repository[entity.Partner], work application.UnitOfWork) *EventService {
	return &EventService{
		eventRepo:   eventRepo,
		partnerRepo: partnerRepo,
		uow:         work,
	}
}

func (e *EventService) List() ([]*entity.Event, error) {
	return e.eventRepo.FindAll()
}

type CreateInput struct {
	Name        string
	Date        time.Time
	PartnerID   string
	Description *string
}

func (e *EventService) FindSections(eventID string) ([]*entity.EventSection, error) {
	event, err := e.eventRepo.FindByID(eventID)
	if err != nil {
		return nil, fmt.Errorf("event service find sections: %v", err)
	}
	var result []*entity.EventSection
	for _, v := range event.Sections.Data {
		result = append(result, v)
	}
	return result, nil
}

func (e *EventService) Create(input CreateInput) (*string, error) {
	partner, err := e.partnerRepo.FindByID(input.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("event service create: %v", err)
	}
	event, err := partner.CreateEvent(entity.PartnerCreateEvent{
		Name:        input.Name,
		Description: input.Description,
		Date:        input.Date,
	})
	if err = e.uow.Begin(); err != nil {
		return nil, fmt.Errorf("event service create: begin event: %v", err)
	}
	err = e.uow.Do(func(u application.UnitOfWork) error {
		eventRepo, errR := u.Repository("EventRepository")
		if errR != nil {
			return fmt.Errorf("event service: do event: %v", errR)
		}
		repo := eventRepo.(domain.Repository[entity.Event])
		if errS := repo.Save(event); errS != nil {
			return fmt.Errorf("event service: do save event: %v", errS)
		}
		return nil
	})
	if err != nil {
		if errRoll := e.uow.Rollback(); errRoll != nil {
			return nil, fmt.Errorf("event service: rollback: %v", errRoll)
		}
		return nil, fmt.Errorf("event service doing: %v", err)
	}
	if err := e.uow.Commit(); err != nil {
		return nil, fmt.Errorf("event service commit: %v", err)
	}
	eventID := event.ID.String()
	return &eventID, nil
}

type AddSectionInput struct {
	Name        string
	Description *string
	TotalSpots  int32
	Price       float64
	EventID     string
}

func (e *EventService) AddSection(input AddSectionInput) error {
	event, err := e.eventRepo.FindByID(input.EventID)
	if err != nil {
		return fmt.Errorf("event service add section find by id: %v", err)
	}
	err = event.AddSection(entity.AddSectionCommand{
		Name:        input.Name,
		Description: input.Description,
		TotalSpots:  input.TotalSpots,
		Price:       input.Price,
	})
	if err != nil {
		return fmt.Errorf("event service add section: %v", err)
	}
	if errB := e.uow.Begin(); errB != nil {
		return fmt.Errorf("event service begin: %v", err)
	}
	if errDo := e.uow.Do(func(u application.UnitOfWork) error {
		repository, errR := u.Repository("EventRepository")
		if errR != nil {
			return fmt.Errorf("event service repository: %v", errR)
		}
		repo := repository.(domain.Repository[entity.Event])
		return repo.Save(event)
	}); errDo != nil {
		if errR := e.uow.Rollback(); errR != nil {
			return fmt.Errorf("event service rollback: %v", errR)
		}
		return fmt.Errorf("event service do: %v", errDo)
	}
	return e.uow.Commit()
}

type ChangeSectionInfo struct {
	EventID           string
	SectionID         string
	Name, Description *string
}

func (e *EventService) ChangeSectionInfo(input ChangeSectionInfo) error {
	event, err := e.eventRepo.FindByID(input.EventID)
	if err != nil {
		return fmt.Errorf("event service change section info: %v", err)
	}
	sectionID, err := domain.NewUUIDFromString(input.SectionID)
	if err != nil {
		return fmt.Errorf("event service change section info: generate uuid %v", err)
	}
	if err = event.ChangeSectionInfo(entity.ChangeSectionInput{
		SectionID:   *sectionID,
		Name:        input.Name,
		Description: input.Description,
	}); err != nil {
		return fmt.Errorf("event service change section info: changing %v", err)
	}
	if err = e.uow.Begin(); err != nil {
		return fmt.Errorf("event service change section info: begin %v", err)
	}
	if err = e.uow.Do(func(u application.UnitOfWork) error {
		repository, errR := u.Repository("EventRepository")
		if errR != nil {
			return errR
		}
		repo := repository.(domain.Repository[entity.Event])
		return repo.Save(event)
	}); err != nil {
		return fmt.Errorf("event service change section info: do %v", err)
	}
	return e.uow.Commit()
}
