package mapvalues

type MapValues[K comparable, V any] struct {
	predicates []func(K, V) V
}

func New[K comparable, V any]() *MapValues[K, V] {
	return &MapValues[K, V]{make([]func(K, V) V, 0)}
}

func (f *MapValues[K, V]) Map(predicate func(K, V) V) *MapValues[K, V] {
	f.predicates = append(f.predicates, predicate)
	return f
}

func (f *MapValues[K, V]) Compose() func(K, V) V {
	return func(key K, value V) V {
		result := value
		for _, predicate := range f.predicates {
			result = predicate(key, result)
		}
		return result
	}
}
