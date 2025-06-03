package db

import (
	"database/sql"
	"errors"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"testing"
	"time"
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
	repo := NewPartnerRepository(database)
	_, err := repo.FindAll()
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
	tx, err := database.Begin()
	if err != nil {
		t.Errorf("expected tx created without error\ngot: %v", err)
		return
	}
	repo := NewPartnerRepository(tx)
	partner, _ := entity.CreatePartner("Danilo Bandeira", time.Now())
	err = repo.Save(partner)
	if err != nil {
		tx.Rollback()
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	tx.Rollback()
}
