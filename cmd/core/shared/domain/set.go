package domain

type Set[T comparable, V any] struct {
	Data map[T]V
}

func NewSet[T comparable, V any]() *Set[T, V] {
	return &Set[T, V]{Data: make(map[T]V)}
}

func (s Set[T, V]) Add(k T, v V) {
	s.Data[k] = v
}

func (s Set[T, V]) Exists(k T) bool {
	_, ok := s.Data[k]
	return ok
}

func (s Set[T, V]) Remove(k T) {
	delete(s.Data, k)
}

func (s Set[T, V]) Size() int {
	return len(s.Data)
}
