package application

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/infra/db"
	"log"
	"testing"
)

func TestCustomerService_List(t *testing.T) {
	conn, _ := db.PostgresConn()
	tx, _ := conn.Begin()
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback tx: %v", err)
		}
	}()
	repo := db.NewCustomerRepository(tx)
	customer1, _ := entity.CreateCustomer("443.376.887-14", "Danilo Bandeira")
	customer2, _ := entity.CreateCustomer("684.705.992-32", "Ana banana")
	customers := []*entity.Customer{customer1, customer2}
	for _, cus := range customers {
		if err := repo.Save(cus); err != nil {
			t.Errorf("error saving: %v\n", err)
			return
		}
	}
	service := NewCustomerService(repo)
	result, err := service.List()
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
	var found bool
	for _, c := range result {
		if c.Equal(customer1) {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected customer 1 to be in list, but it's not")
		return
	}
	found = false
	for _, c := range result {
		if c.Equal(customer2) {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected customer 2 to be in list, but it's not")
		return
	}
}

func TestCustomerService_Register(t *testing.T) {
	conn, _ := db.PostgresConn()
	tx, _ := conn.Begin()
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback tx: %v", err)
		}
	}()
	repo := db.NewCustomerRepository(tx)
	service := NewCustomerService(repo)
	err := service.Register(RegisterInput{
		Name: "Danilo Bandeira",
		CPF:  "148.794.614-74",
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
}
