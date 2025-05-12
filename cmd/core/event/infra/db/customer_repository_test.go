package db

import (
	"database/sql"
	"errors"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"testing"
)

func TestCustomerRepository_FindAll(t *testing.T) {
	database, _ := PostgresConn()
	repo := NewCustomerRepository(database)
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

func TestCustomerRepository_Save(t *testing.T) {
	database, _ := PostgresConn()
	tx, err := database.Begin()
	if err != nil {
		t.Errorf("expected tx created without error\ngot: %v", err)
		return
	}
	repo := NewCustomerRepository(tx)
	customer, _ := entity.CreateCustomer("141.053.121-03", "Danilo Bandeira")
	err = repo.Save(customer)
	if err != nil {
		tx.Rollback()
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	tx.Rollback()
}
