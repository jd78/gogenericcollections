package arraycomposition

import "github.com/jd78/gogenericcollections/composition"

type Filter[K comparable] func(K) bool

func (f Filter[K]) GetType() composition.FunctionType {
	return composition.Filter
}

func (f Filter[K]) Exec(val K) bool {
	return f(val)
}

// func (f *filter[K]) compose() func(K) bool {
// 	return func(value K) bool {
// 		for _, predicate := range f.predicates {
// 			if !predicate(value) {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// }
