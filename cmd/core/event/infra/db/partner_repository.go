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
	rows, err := p.db.QueryContext(ctx, "select id, name from partners")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("find all: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("find all: finding: %w", err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			err = fmt.Errorf("find all: close: %v", err)
		}
	}()
	for rows.Next() {
		var partner entity.Partner
		errScan := rows.Scan(&partner.ID, &partner.Name)
		if errScan != nil {
			return nil, fmt.Errorf("find all: scanning: %w", err)
		}
		result = append(result, &partner)
	}
	return result, err
}

func NewRepository(db *sql.DB) *PartnerRepository {
	return &PartnerRepository{db}
}
