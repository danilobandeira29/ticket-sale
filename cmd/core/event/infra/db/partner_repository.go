package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type PartnerRepository struct {
	executor Executor
}

func (p *PartnerRepository) FindAll() (result []*entity.Partner, err error) {
	rows, err := p.executor.Query("select id, name from partners;")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("partner repository: find all: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("partner repository: find all: finding: %w", err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			err = fmt.Errorf("partner repository: find all: close: %v", err)
		}
	}()
	for rows.Next() {
		var partner entity.Partner
		errScan := rows.Scan(&partner.ID, &partner.Name)
		if errScan != nil {
			return nil, fmt.Errorf("partner repository: find all: scanning: %w", errScan)
		}
		result = append(result, &partner)
	}
	return result, err
}

func (p *PartnerRepository) Save(partner *entity.Partner) error {
	_, err := p.executor.Exec("insert into partners(id, name) values ($1, $2);", partner.ID.String(), partner.Name)
	if err != nil {
		return fmt.Errorf("partner repository: exec %v", err)
	}
	return nil
}

func (p *PartnerRepository) FindByID(id any) (*entity.Partner, error) {
	partnerID, ok := id.(string)
	if !ok {
		return nil, fmt.Errorf("partner repository: find by id: invalid id: %T", id)
	}
	row := p.executor.QueryRow("select id, name from partners where id = $1", partnerID)
	var partner entity.Partner
	if err := row.Scan(&partner.ID, &partner.Name); err != nil {
		return nil, fmt.Errorf("partner repository find by id: scanning: %v", err)
	}
	return &partner, nil
}

func (p *PartnerRepository) Delete(id any) error {
	// todo: need to use this id
	_, ok := id.(entity.PartnerID)
	if !ok {
		return fmt.Errorf("partnert repository: delete: invalid id: %T", id)
	}
	return nil
}

func NewPartnerRepository(exec Executor) domain.Repository[entity.Partner] {
	return &PartnerRepository{exec}
}
