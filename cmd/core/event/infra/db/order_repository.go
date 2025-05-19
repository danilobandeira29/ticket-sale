package db

import (
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/domain/entity"
)

type OrderRepository struct{}

func (o *OrderRepository) Save(t *entity.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepository) FindByID(id any) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepository) FindAll() ([]*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepository) Delete(id any) error {
	//TODO implement me
	panic("implement me")
}
