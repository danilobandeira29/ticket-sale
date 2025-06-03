package application

import (
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"sync"
	"time"
)

type OrderService struct {
	orderRepo           domain.Repository[entity.Order]
	customerRepo        domain.Repository[entity.Customer]
	spotReservationRepo domain.Repository[entity.SpotReservation]
	event               domain.Repository[entity.Event]
	uow                 application.UnitOfWork
}

func NewOrderService(orderRepo domain.Repository[entity.Order], uow application.UnitOfWork) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		uow:       uow,
	}
}

type OrderCreateInput struct {
	EventID, SectionID, SpotID, CustomerID string
}

// TODO: quando a ordem de serviço falhar, mas não por motivo de reserva de lugar, então deve-se guardar no banco essa order
func (o *OrderService) Create(input OrderCreateInput) (*entity.Order, error) {
	var (
		customer *entity.Customer
		event    *entity.Event
		errCh    = make(chan error, 2)
		wg       sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		c, err := o.customerRepo.FindByID(input.CustomerID)
		if err != nil {
			errCh <- fmt.Errorf("order service create: find customer: %v", err)
			return
		}
		customer = c
		errCh <- nil
	}()
	go func() {
		defer wg.Done()
		e, err := o.event.FindByID(input.EventID)
		if err != nil {
			errCh <- fmt.Errorf("order return create: find event: %v", err)
			return
		}
		event = e
		errCh <- nil
	}()
	wg.Wait()
	close(errCh)
	for errChannel := range errCh {
		if errChannel != nil {
			return nil, errChannel
		}
	}
	sectionID, err := domain.NewUUIDFromString(input.SectionID)
	if err != nil {
		return nil, fmt.Errorf("order service create uuid: %v", err)
	}
	spotID, err := domain.NewUUIDFromString(input.SpotID)
	if err != nil {
		return nil, fmt.Errorf("order service create uuid spotid: %v", err)
	}
	allowReserveSpot, err := event.AllowReserveSpot(entity.AllowReserveSpotInput{
		SectionID: *sectionID,
		SpotID:    *spotID,
	})
	if err != nil {
		return nil, fmt.Errorf("order service: allow: %v", err)
	}
	if !allowReserveSpot {
		return nil, fmt.Errorf("order service: not allow reserve spot")
	}
	reservation, err := o.spotReservationRepo.FindByID(spotID)
	if err != nil {
		return nil, fmt.Errorf("order service: find reservation: %v", err)
	}
	if reservation != nil {
		return nil, errors.New("order service: already reserved")
	}
	spotReservation := entity.NewSpotReservation(*spotID, time.Now(), customer.ID)
	section, err := event.Section(*sectionID)
	if err != nil {
		return nil, fmt.Errorf("order service: %w", err)
	}
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("order service: order id: %w", err)
	}
	order := entity.NewOrder(entity.OrderProps{
		ID:          *id,
		CustomerID:  customer.ID,
		Amount:      section.Price,
		EventSpotID: *spotID,
	})
	if errB := o.uow.Begin(); errB != nil {
		return nil, errB
	}
	err = o.uow.Do(func(u application.UnitOfWork) error {
		repo, errR := u.Repository("SpotReservationRepository")
		if errR != nil {
			return errR
		}
		spotRepo := repo.(domain.Repository[entity.SpotReservation])
		if errS := spotRepo.Save(spotReservation); err != nil {
			return fmt.Errorf("order service: saving reservation: %v", errS)
		}
		repo, errR = u.Repository("OrderRepository")
		if errR != nil {
			return errR
		}
		orderRepo := repo.(domain.Repository[entity.Order])
		if errS := orderRepo.Save(order); errS != nil {
			return fmt.Errorf("order service: saving order: %v", errS)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if errC := o.uow.Commit(); errC != nil {
		return nil, errC
	}
	return order, nil
}
