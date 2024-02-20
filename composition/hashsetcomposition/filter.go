package hashsetcomposition

import "github.com/jd78/gogenericcollections/composition"

type Filter[K comparable] func(K) bool

func (f Filter[K]) GetType() composition.FunctionType {
	return composition.Filter
}

func (f Filter[K]) Exec(key K) bool {
	return f(key)
}
