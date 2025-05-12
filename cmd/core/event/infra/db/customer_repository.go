package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/domain"
)

type CustomerRepository struct {
	executor Executor
}

func (c CustomerRepository) Save(customer *entity.Customer) error {
	_, err := c.executor.Exec("insert into customers(id, name, cpf) values ($1, $2, $3);", customer.ID.String(), customer.Name.String(), customer.CPF.Value())
	if err != nil {
		return fmt.Errorf("customer repository: exec %v", err)
	}
	return nil
}

func (c CustomerRepository) FindByID(id any) (*entity.Customer, error) {
	// todo: need to use this id
	_, ok := id.(entity.CustomerID)
	if !ok {
		return nil, fmt.Errorf("customer repository: find by id: invalid id: %T", id)
	}
	return nil, nil
}

func (c CustomerRepository) FindAll() (result []*entity.Customer, err error) {
	rows, err := c.executor.Query("select id, name, cpf from customers;")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("customer repository: find all: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("customer repository: find all: finding: %w", err)
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			err = fmt.Errorf("customer repository: find all: close: %v", err)
		}
	}()
	for rows.Next() {
		var customer entity.Customer
		errScan := rows.Scan(&customer.ID, &customer.Name, &customer.CPF)
		if errScan != nil {
			return nil, fmt.Errorf("customer repository: find all: scanning: %w", errScan)
		}
		result = append(result, &customer)
	}
	return result, err
}

func (c CustomerRepository) Delete(id any) error {
	// todo: need to use this id
	_, ok := id.(entity.CustomerID)
	if !ok {
		return fmt.Errorf("customer repository: delete: invalid id: %T", id)
	}
	return nil
}

func NewCustomerRepository(executor Executor) domain.Repository[entity.Customer] {
	return &CustomerRepository{executor: executor}
}
