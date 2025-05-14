package entity

import (
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type CustomerID = domain.UUID

type Customer struct {
	aggregate domain.AggregateRoot
	ID        CustomerID
	CPF       *domain.CPF
	Name      *domain.Name
}

func (c *Customer) String() string {
	return c.aggregate.String(c)
}

func NewCustomer(id, cpf, name string) (*Customer, error) {
	c, err := domain.NewCPF(cpf)
	if err != nil {
		return nil, fmt.Errorf("new customer: creating cpf %v", err)
	}
	i, err := domain.NewUUIDFromString(id)
	if err != nil {
		return nil, fmt.Errorf("new customer: creating id %v", err)
	}
	return &Customer{
		ID:   *i,
		CPF:  c,
		Name: domain.NewName(name),
	}, nil
}

func CreateCustomer(cpf, name string) (*Customer, error) {
	id, err := domain.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("customer entity create: %v", err)
	}
	c, err := domain.NewCPF(cpf)
	if err != nil {
		return nil, fmt.Errorf("new customer: creating cpf %v", err)
	}
	return &Customer{
		ID:   *id,
		CPF:  c,
		Name: domain.NewName(name),
	}, nil
}

func (c *Customer) Equal(o *Customer) bool {
	return c.ID.Equal(o.ID) &&
		c.CPF.Equal(o.CPF) &&
		c.Name.Equal(o.Name)
}
