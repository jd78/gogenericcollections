package mapcomposition

import "github.com/jd78/gogenericcollections/composition"

type MapValues[K comparable, V any] func(K, V) V

func (f MapValues[K, V]) GetType() composition.FunctionType {
	return composition.MapValues
}

func (f MapValues[K, V]) Exec(key K, value V) V {
	return f(key, value)
}
