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
	eventID, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("event repository: find by id: invalid id: %T", id)
	}
	var event entity.Event
	err := p.executor.QueryRow(`
		SELECT id, name, description, date, is_published, total_spots, total_spots_reserved, partner_id
		FROM events WHERE id = $1
	`, eventID).Scan(&event.ID, &event.Name, &event.Description, &event.Date, &event.IsPublished, &event.TotalSpots, &event.TotalSpotsReserved, &event.PartnerID)
	if err != nil {
		return nil, fmt.Errorf("event repository: find by id queryrow: %v", err)
	}
	rows, err := p.executor.Query(`
		SELECT id, name, description, is_published, total_spots, total_spots_reserved, price
		FROM event_sections WHERE event_id = $1
	`, eventID)
	if err != nil {
		return nil, fmt.Errorf("event repository: find sections: %v", err)
	}
	defer rows.Close()
	sections := domain.NewSet[string, *entity.EventSection]()
	for rows.Next() {
		var section entity.EventSection
		err = rows.Scan(&section.ID, &section.Name, &section.Description, &section.IsPublished, &section.TotalSpots, &section.TotalSpotsReserved, &section.Price)
		if err != nil {
			return nil, fmt.Errorf("event repository: scan section: %v", err)
		}
		spotsRows, err := p.executor.Query(`
			SELECT id, location, is_published, is_reserved
			FROM event_spots
			WHERE event_section_id = $1
		`, section.ID.String())
		if err != nil {
			return nil, fmt.Errorf("event repository: find spots: %v", err)
		}
		eventSpotSet := domain.NewSet[string, *entity.EventSpot]()
		for spotsRows.Next() {
			var spot entity.EventSpot
			err = spotsRows.Scan(&spot.ID, &spot.Location, &spot.IsPublished, &spot.IsReserved)
			if err != nil {
				spotsRows.Close()
				return nil, fmt.Errorf("event repository: scan spot: %v", err)
			}
			eventSpot, err := entity.NewEventSpot(entity.EventSpotProps{
				ID:          &spot.ID,
				Location:    spot.Location,
				IsReserved:  spot.IsReserved,
				IsPublished: spot.IsPublished,
			})
			if err != nil {
				return nil, fmt.Errorf("event repository: new event spot: %v", err)
			}
			eventSpotSet.Add(spot.ID.String(), eventSpot)
		}
		spotsRows.Close()
		sec, err := entity.NewEventSection(entity.EventSectionProps{
			ID:                 &section.ID,
			Name:               section.Name,
			Description:        section.Description,
			TotalSpots:         section.TotalSpots,
			TotalSpotsReserved: section.TotalSpotsReserved,
			IsPublished:        section.IsPublished,
			Price:              section.Price,
			Spots:              *eventSpotSet,
		})
		if err != nil {
			return nil, fmt.Errorf("event repository: create section entity: %v", err)
		}
		sections.Add(section.ID.String(), sec)
	}
	entityEvent, err := entity.NewEvent(entity.EventProps{
		ID:                 &event.ID,
		Name:               event.Name,
		Description:        event.Description,
		Date:               event.Date,
		IsPublished:        event.IsPublished,
		TotalSpots:         event.TotalSpots,
		TotalSpotsReserved: event.TotalSpotsReserved,
		PartnerID:          &event.PartnerID,
		EventSectionSet:    *sections,
	})
	if err != nil {
		return nil, fmt.Errorf("event repository: create event entity: %v", err)
	}
	return entityEvent, nil
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
