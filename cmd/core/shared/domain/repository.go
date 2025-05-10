package domain

type Repository[T any] interface {
	Save(t *T) error
	FindByID(id any) (*T, error)
	FindAll() ([]*T, error)
	Delete(id any) error
}
