package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
)

type PartnerRepository struct {
	db *sql.DB
}

func (p *PartnerRepository) FindAll(ctx context.Context) (result []*entity.Partner, err error) {
	rows, err := p.db.QueryContext(ctx, "select id, name from partners;")
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

func (p *PartnerRepository) Save(partner entity.Partner) error {
	_, err := p.db.Exec("insert into partners(id, name) values ($1, $2);", partner.ID.String(), partner.Name)
	if err != nil {
		return fmt.Errorf("partner repository: exec %v", err)
	}
	return nil
}

func NewRepository(db *sql.DB) *PartnerRepository {
	return &PartnerRepository{db}
}
