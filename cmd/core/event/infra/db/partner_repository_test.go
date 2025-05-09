package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"testing"
)

func TestPostgresConn(t *testing.T) {
	_, err := PostgresConn()
	if err != nil {
		t.Errorf("error conn: %v", err)
		return
	}
}

func TestPartnerRepository_FindAll(t *testing.T) {
	database, _ := PostgresConn()
	repo := NewRepository(database)
	_, err := repo.FindAll(context.Background())
	if errors.Is(err, sql.ErrNoRows) {
		t.Errorf("no rows")
		return
	}
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
}

func TestPartnerRepository_Save(t *testing.T) {
	database, _ := PostgresConn()
	repo := NewRepository(database)
	partner, _ := entity.CreatePartner("Danilo Bandeira")
	err := repo.Save(*partner)
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
}
