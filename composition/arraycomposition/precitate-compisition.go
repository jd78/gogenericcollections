package arraycomposition

import "github.com/jd78/gogenericcollections/composition"

type functionType interface {
	GetType() composition.FunctionType
}

type ArrayComposition[K comparable] struct {
	predicates []functionType
}

func New[K comparable]() *ArrayComposition[K] {
	return &ArrayComposition[K]{[]functionType{}}
}

func (a *ArrayComposition[K]) AddFunction(fn functionType) *ArrayComposition[K] {
	a.predicates = append(a.predicates, fn)
	return a
}

func (a *ArrayComposition[K]) GetPredicates() []functionType {
	return a.predicates
}
