package genericmap

import (
	mapfilter "github.com/jd78/gogenericcollections/composition/filter/map-filter"
	mapvalues "github.com/jd78/gogenericcollections/composition/transform/map-values"
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
func (m GenericMap[K, V]) Filter(predicate func(K, V) bool) *ProxedMap[K, V] {
	filter := mapfilter.New[K, V]()
	filter.AddFilter(predicate)
	return NewWithFilter(m, filter)
}

// MapValues applies the predicate to the values
func (m GenericMap[K, V]) MapValues(predicate func(K, V) V) *ProxedMap[K, V] {
	mapValues := mapvalues.New[K, V]()
	mapValues.Map(predicate)
	return NewWithMapValues[K, V](m, mapValues)
}

// AddAll add all elements contained in the passed map
func (m GenericMap[K, V]) AddAll(origin map[K]V) GenericMap[K, V] {
	for k, v := range origin {
		m.Add(k, v)
	}
	return m
}
