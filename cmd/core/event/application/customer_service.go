package application

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type CustomerService struct {
	repo domain.Repository[entity.Customer]
}

func NewCustomerService(repo domain.Repository[entity.Customer]) *CustomerService {
	return &CustomerService{repo}
}

func (c *CustomerService) List() ([]*entity.Customer, error) {
	return c.repo.FindAll()
}

type RegisterInput struct {
	Name string
	CPF  string
}

func (c *CustomerService) Register(input RegisterInput) error {
	customer, err := entity.CreateCustomer(input.CPF, input.Name)
	if err != nil {
		return fmt.Errorf("customer service: creating: %v", err)
	}
	err = c.repo.Save(customer)
	if err != nil {
		return fmt.Errorf("customer service: saving: %v", err)
	}
	return nil
}
