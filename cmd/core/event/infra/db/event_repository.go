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
	_, err := p.executor.Exec(`
		insert into events(
			id, 
			name, 
			description, 
			date, 
			is_published, 
			partner_id, 
			total_spots, 
			total_spots_reserved
		) values ($1, $2, $3, $4, $5, $6, $7, $8);`, event.ID.String(), event.Name, event.Description, event.Date, event.IsPublished, event.PartnerID.String(), event.TotalSpots, event.TotalSpotsReserved)
	if err != nil {
		return fmt.Errorf("event repository: exec %v", err)
	}
	for _, section := range event.Sections.Data {
		_, errSection := p.executor.Exec(`
			insert into event_sections(
				id,
				name,
				description,
				is_published,
				total_spots,
				total_spots_reserved,
				price,
				event_id
			) values ($1, $2, $3, $4, $5, $6, $7, $8);`,
			section.ID.String(),
			section.Name,
			section.Description,
			section.IsPublished,
			section.TotalSpots,
			section.TotalSpotsReserved,
			section.Price,
			event.ID.String(),
		)
		if errSection != nil {
			return fmt.Errorf("event repository: save section: %w", errSection)
		}
		for _, spot := range section.Spots.Data {
			_, errSpot := p.executor.Exec(`
				insert into event_spots(
					id,
					event_section_id,
					location,
					is_published,
					is_reserved
				) values ($1, $2, $3, $4, $5);`,
				spot.ID.String(),
				section.ID.String(),
				spot.Location,
				spot.IsPublished,
				spot.IsReserved,
			)
			if errSpot != nil {
				return fmt.Errorf("event repository: save spot: %w", errSpot)
			}
		}
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
