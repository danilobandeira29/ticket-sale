package application

import "github.com/danilobandeira29/ticket-sale/cmd/core/event/infra/db"

type UnitOfWork interface {
	Begin() error
	Do(fn func(u UnitOfWork) error) error
	RegisterFactory(n string, fn func(exec db.Executor) any)
	Repository(n string) (any, error)
	Commit() error
	Rollback() error
}
