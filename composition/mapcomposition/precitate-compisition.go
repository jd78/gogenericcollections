package mapcomposition

import "github.com/jd78/gogenericcollections/composition"

type functionType interface {
	GetType() composition.FunctionType
}

type MapComposition[K comparable, V any] struct {
	predicates []functionType
}

func New[K comparable, V any]() *MapComposition[K, V] {
	return &MapComposition[K, V]{[]functionType{}}
}

func (a *MapComposition[K, V]) AddFunction(fn functionType) *MapComposition[K, V] {
	a.predicates = append(a.predicates, fn)
	return a
}

func (a *MapComposition[K, V]) GetPredicates() []functionType {
	return a.predicates
}
