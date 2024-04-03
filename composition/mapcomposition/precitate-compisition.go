package mapcomposition

import "github.com/jd78/gogenericcollections/composition"

type functionType interface {
	GetType() composition.FunctionType
}

type MapComposition[K comparable, V any] struct {
	predicates []functionType
	limit      int
}

func New[K comparable, V any]() *MapComposition[K, V] {
	return &MapComposition[K, V]{[]functionType{}, 0}
}

func (a *MapComposition[K, V]) AddFunction(fn functionType) *MapComposition[K, V] {
	a.predicates = append(a.predicates, fn)
	return a
}

func (a *MapComposition[K, V]) AddLimit(limit int) *MapComposition[K, V] {
	a.limit = limit
	return a
}

func (a *MapComposition[K, V]) GetPredicates() []functionType {
	return a.predicates
}

func (a *MapComposition[K, V]) GetLimit() int {
	return a.limit
}
