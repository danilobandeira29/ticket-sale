package unitofwork

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/infra/db"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
	"testing"
)

func TestUow_Do(t *testing.T) {
	conn, _ := db.PostgresConn()
	worker := NewUoW(conn)
	repo := db.NewCustomerRepository(conn)
	// todo: fix this conn, repo uses conn and worker user tx. both needs to use tx
	worker.Register("CustomerRepository", repo)
	_ = worker.Begin()
	customer, _ := entity.CreateCustomer("824.622.083-72", "Danilo Bandeira")
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
	worker.Rollback()
}
