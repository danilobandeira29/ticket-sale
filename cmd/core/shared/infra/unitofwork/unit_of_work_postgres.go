package unitofwork

import (
	"database/sql"
	"fmt"
	"github.com/danilobandeira29/ticket-sale/cmd/core/shared/application"
)

type uow struct {
	db   *sql.DB
	tx   *sql.Tx
	repo map[string]any
}

func NewUoW(db *sql.DB) application.UnitOfWork {
	return &uow{
		db:   db,
		tx:   nil,
		repo: make(map[string]any),
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

func (u *uow) Do(fn func(tx application.UnitOfWork) error) error {
	if u.tx == nil {
		return fmt.Errorf("uow: must call Begin() before")
	}
	err := fn(u)
	if err != nil {
		return fmt.Errorf("uow: rollback error: %v", err)
	}
	return nil
}

func (u *uow) Register(n string, repository any) {
	u.repo[n] = repository
}

func (u *uow) Repository(n string) (any, error) {
	repo, ok := u.repo[n]
	if !ok {
		return nil, fmt.Errorf("uowtx: repository not found")
	}
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
