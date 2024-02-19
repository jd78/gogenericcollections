package mapcomposition

import "github.com/jd78/gogenericcollections/composition"

type Filter[K comparable, V any] func(K, V) bool

func (f Filter[K, V]) GetType() composition.FunctionType {
	return composition.Filter
}

func (f Filter[K, V]) Exec(key K, value V) bool {
	return f(key, value)
}
