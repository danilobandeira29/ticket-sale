package unitofwork

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/event/infra/db"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
)

type repoFactory func(executor db.Executor) any

type uow struct {
	db          *sql.DB
	tx          *sql.Tx
	repo        map[string]any
	repoFactory map[string]repoFactory
}

func NewUoW(db *sql.DB) application.UnitOfWork {
	return &uow{
		db:          db,
		tx:          nil,
		repo:        make(map[string]any),
		repoFactory: make(map[string]repoFactory),
	}
}

func (u *uow) Begin() error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	u.tx = tx
	return nil
}

func (u *uow) Do(fn func(_ application.UnitOfWork) error) error {
	if u.tx == nil {
		return fmt.Errorf("uow: must call Begin() before")
	}
	err := fn(u)
	if err != nil {
		return fmt.Errorf("uow: rollback error: %v", err)
	}
	return nil
}

func (u *uow) RegisterFactory(n string, fac func(exec db.Executor) any) {
	u.repoFactory[n] = fac
}

func (u *uow) Repository(n string) (any, error) {
	if repo, ok := u.repo[n]; ok {
		return repo, nil
	}
	factory, ok := u.repoFactory[n]
	if !ok {
		return nil, errors.New("uow: repository factory not found")
	}
	repo := factory(u.tx)
	u.repo[n] = repo
	return repo, nil
}

func (u *uow) Commit() error {
	defer func() { u.tx = nil }()
	return u.tx.Commit()
}

func (u *uow) Rollback() error {
	defer func() { u.tx = nil }()
	return u.tx.Rollback()
}
