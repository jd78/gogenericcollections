package arrayfilter

type Filter[K comparable] struct {
	predicates []func(K) bool
}

func New[K comparable]() *Filter[K] {
	return &Filter[K]{make([]func(K) bool, 0)}
}

func (f *Filter[K]) AddFilter(filter func(K) bool) *Filter[K] {
	f.predicates = append(f.predicates, filter)
	return f
}

func (f *Filter[K]) Compose() func(K) bool {
	return func(value K) bool {
		for _, predicate := range f.predicates {
			if !predicate(value) {
				return false
			}
		}
		return true
	}
}
