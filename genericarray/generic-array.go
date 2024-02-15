package genericarray

import "github.com/jd78/gogenericcollections/genericmap"

type GenericArray[K comparable] []K

// New initialize a generic array
func New[K comparable]() GenericArray[K] {
	return make(GenericArray[K], 0)
}

// From creates a generic array from a passed array
func From[K comparable](origin []K) GenericArray[K] {
	a := make(GenericArray[K], 0)
	for _, v := range origin {
		a.Add(v)
	}
	return a
}

// Add adds an element
func (a *GenericArray[K]) Add(val K) {
	*a = append(*a, val)
}

// AddAll adds all elements of the passed array
func (a *GenericArray[K]) AddAll(origin []K) GenericArray[K] {
	for _, v := range origin {
		a.Add(v)
	}
	return *a
}

// MapValues transforms the array values running the passed predicate
func (a GenericArray[K]) MapValues(predicate func(K) K) GenericArray[K] {
	mapped := New[K]()
	for _, v := range a {
		mapped.Add(predicate(v))
	}
	return mapped
}

// MapArray creates a new array of type V running the passed predicate
func MapArray[K, V comparable](a GenericArray[K], predicate func(K) V) GenericArray[V] {
	mappedArray := New[V]()

	for _, item := range a {
		mappedItem := predicate(item)
		mappedArray.Add(mappedItem)
	}

	return mappedArray
}

// ArrayToMap creates a new map of type V, Z running the passed predicate
func ArrayToMap[K, V, Z comparable](a GenericArray[K], predicate func(K) (V, Z)) genericmap.GenericMap[V, Z] {
	m := genericmap.New[V, Z]()

	for _, item := range a {
		first, second := predicate(item)
		m.Add(first, second)
	}

	return m
}
