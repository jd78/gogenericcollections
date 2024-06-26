package genericmap

import (
	"github.com/jd78/gogenericcollections/composition/mapcomposition"
)

type GenericMap[K comparable, V any] map[K]V

// New initialize a new generic map
func New[K comparable, V any]() GenericMap[K, V] {
	return make(GenericMap[K, V])
}

// HasKey checks if the map has the passed key
func (m GenericMap[K, V]) HasKey(key K) bool {
	_, exists := m[key]
	return exists
}

// Add adds an element
func (m GenericMap[K, V]) Add(key K, value V) GenericMap[K, V] {
	m[key] = value
	return m
}

// Get gets a value
func (m GenericMap[K, V]) Get(key K) V {
	return m[key]
}

// Delete deletes a key
func (m GenericMap[K, V]) Delete(key K) {
	if !m.HasKey(key) {
		return
	}
	var zero V
	m[key] = zero
	delete(m, key)
}

// Filter proxy filtered map by predicates
func (m GenericMap[K, V]) Filter(predicate mapcomposition.Filter[K, V]) *ProxedMap[K, V] {
	ac := mapcomposition.New[K, V]()
	ac.AddFunction(predicate)
	pa := newProxedMap(m, ac)
	return pa
}

// Limit the result
func (m GenericMap[K, V]) Limit(limit int) *ProxedMap[K, V] {
	ac := mapcomposition.New[K, V]()
	ac.AddLimit(limit)
	pa := newProxedMap(m, ac)
	return pa
}

// MapValues applies the predicate to the values
func (m GenericMap[K, V]) MapValues(predicate mapcomposition.MapValues[K, V]) *ProxedMap[K, V] {
	ac := mapcomposition.New[K, V]()
	ac.AddFunction(predicate)
	pa := newProxedMap(m, ac)
	return pa
}

// AddAll add all elements contained in the passed map
func (m GenericMap[K, V]) AddAll(origin map[K]V) GenericMap[K, V] {
	for k, v := range origin {
		m.Add(k, v)
	}
	return m
}

func ToList[K comparable, V, Z any](m GenericMap[K, V], selector func(K, V) Z) []Z {
	l := make([]Z, len(m))
	index := 0
	for k, v := range m {
		l[index] = selector(k, v)
		index++
	}
	return l
}
