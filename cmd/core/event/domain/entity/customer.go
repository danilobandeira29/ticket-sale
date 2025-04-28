package entity

import (
	"fmt"
	"github.com/danilobandeira29/ddd-go/cmd/core/shared/domain"
	"github.com/google/uuid"
)

type Customer struct {
	domain.AggregateRoot
	ID   string
	CPF  string
	Name string
}

func NewCustomer(id, cpf, name string) (*Customer, error) {
	return &Customer{
		ID:   id,
		CPF:  cpf,
		Name: name,
	}, nil
}

func CreateCustomer(cpf, name string) (*Customer, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("customer entity create: %v", err)
	}
	return &Customer{
		ID:   id.String(),
		CPF:  cpf,
		Name: name,
	}, nil
}

func (c *Customer) String() string {
	return c.AggregateRoot.String(c)
}
