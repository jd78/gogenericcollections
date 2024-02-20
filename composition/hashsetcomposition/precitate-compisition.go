package hashsetcomposition

import "github.com/jd78/gogenericcollections/composition"

type functionType interface {
	GetType() composition.FunctionType
}

type HashsetComposition[K comparable] struct {
	predicates []functionType
}

func New[K comparable]() *HashsetComposition[K] {
	return &HashsetComposition[K]{[]functionType{}}
}

func (a *HashsetComposition[K]) AddFunction(fn functionType) *HashsetComposition[K] {
	a.predicates = append(a.predicates, fn)
	return a
}

func (a *HashsetComposition[K]) GetPredicates() []functionType {
	return a.predicates
}
