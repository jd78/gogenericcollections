package arraycomposition

import "github.com/jd78/gogenericcollections/composition"

type MapValues[K comparable] func(K) K

func (f MapValues[K]) GetType() composition.FunctionType {
	return composition.MapValues
}

func (f MapValues[K]) Exec(val K) K {
	return f(val)
}
