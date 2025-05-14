package application

type UnitOfWork interface {
	Begin() error
	Do(fn func(u UnitOfWork) error) error
	Register(n string, repository any)
	Repository(n string) (any, error)
	Commit() error
	Rollback() error
}
