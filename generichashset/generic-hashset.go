package generichashset

import (
	"github.com/jd78/gogenericcollections/composition/hashsetcomposition"
)

type GenericHashset[K comparable] map[K]struct{}

// New initialize a new generic hashset
func New[K comparable]() GenericHashset[K] {
	return make(GenericHashset[K])
}

// HasKey checks if the hashset has the passed key
func (m GenericHashset[K]) HasKey(key K) bool {
	_, exists := m[key]
	return exists
}

// Add adds an element
func (m GenericHashset[K]) Add(key K) GenericHashset[K] {
	m[key] = struct{}{}
	return m
}

// Delete deletes a key
func (m GenericHashset[K]) Delete(key K) {
	if !m.HasKey(key) {
		return
	}
	var zero struct{}
	m[key] = zero
	delete(m, key)
}

// Filter proxy filtered hashset by predicates
func (m GenericHashset[K]) Filter(predicate hashsetcomposition.Filter[K]) *ProxedHashset[K] {
	ac := hashsetcomposition.New[K]()
	ac.AddFunction(predicate)
	pa := newProxedHashset(m, ac)
	return pa
}

// AddAll add all elements contained in the passed map
func (m GenericHashset[K]) AddAll(origin map[K]struct{}) GenericHashset[K] {
	for k, _ := range origin {
		m.Add(k)
	}
	return m
}
