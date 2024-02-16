package filter

type Filter[K comparable, V any] struct {
	predicates []func(K, V) bool
}

func New[K comparable, V any]() *Filter[K, V] {
	return &Filter[K, V]{make([]func(K, V) bool, 0)}
}

func (f *Filter[K, V]) AddFilter(filter func(K, V) bool) *Filter[K, V] {
	f.predicates = append(f.predicates, filter)
	return f
}

func (f *Filter[K, V]) Compose() func(K, V) bool {
	return func(key K, value V) bool {
		for _, predicate := range f.predicates {
			if !predicate(key, value) {
				return false
			}
		}
		return true
	}
}
