package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type EventRepository struct {
	executor Executor
}

func (p *EventRepository) FindAll() (result []*entity.Event, err error) {
	rows, err := p.executor.Query("select id, name from events;")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("event repository: find all: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("event repository: find all: finding: %w", err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			err = fmt.Errorf("event repository: find all: close: %v", err)
		}
	}()
	for rows.Next() {
		var event entity.Event
		errScan := rows.Scan(&event.ID, &event.Name)
		if errScan != nil {
			return nil, fmt.Errorf("event repository: find all: scanning: %w", errScan)
		}
		result = append(result, &event)
	}
	return result, err
}

func (p *EventRepository) Save(event *entity.Event) error {
	_, err := p.executor.Exec("insert into events(id, name) values ($1, $2);", event.ID.String(), event.Name)
	if err != nil {
		return fmt.Errorf("event repository: exec %v", err)
	}
	return nil
}

func (p *EventRepository) FindByID(id any) (*entity.Event, error) {
	// todo: need to use this id
	_, ok := id.(entity.EventID)
	if !ok {
		return nil, fmt.Errorf("event repository: find by id: invalid id: %T", id)
	}
	return nil, nil
}

func (p *EventRepository) Delete(id any) error {
	// todo: need to use this id
	_, ok := id.(entity.EventID)
	if !ok {
		return fmt.Errorf("event repository: delete: invalid id: %T", id)
	}
	return nil
}

func NewEventRepository(exec Executor) domain.Repository[entity.Event] {
	return &EventRepository{exec}
}
