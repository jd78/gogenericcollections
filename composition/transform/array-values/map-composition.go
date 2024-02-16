package arrayvalues

type MapValues[K comparable] struct {
	predicates []func(K) K
}

func New[K comparable]() *MapValues[K] {
	return &MapValues[K]{make([]func(K) K, 0)}
}

func (f *MapValues[K]) Map(predicate func(K) K) *MapValues[K] {
	f.predicates = append(f.predicates, predicate)
	return f
}

func (f *MapValues[K]) Compose() func(K) K {
	return func(value K) K {
		result := value
		for _, predicate := range f.predicates {
			result = predicate(value)
		}
		return result
	}
}
