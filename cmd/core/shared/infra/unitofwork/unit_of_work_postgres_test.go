package unitofwork

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/infra/db"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"log"
	"testing"
)

func TestUow_Do(t *testing.T) {
	conn, _ := db.PostgresConn()
	worker := NewUoW(conn)
	_ = worker.Begin()
	defer func() {
		if err := worker.Rollback(); err != nil {
			log.Printf("error rollback worker: %v", err)
		}
	}()
	worker.RegisterFactory("CustomerRepository", func(exec db.Executor) any {
		return db.NewCustomerRepository(exec)
	})
	customer, _ := entity.CreateCustomer("586.285.965-93", "Danilo Bandeira")
	err := worker.Do(func(u application.UnitOfWork) error {
		r, err := u.Repository("CustomerRepository")
		repoCustomer := r.(domain.Repository[entity.Customer])
		if err != nil {
			return err
		}
		return repoCustomer.Save(customer)
	})
	if err != nil {
		t.Errorf("expected error to be empty\ngot: %v", err)
		return
	}
}
