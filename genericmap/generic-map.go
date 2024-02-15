package genericmap

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
func (m GenericMap[K, V]) Add(key K, value V) {
	m[key] = value
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

// Filter filters the map creating a new map using the passing predicate
func (m GenericMap[K, V]) Filter(predicate func(K, V) bool) GenericMap[K, V] {
	filteredMap := New[K, V]()
	for key, value := range m {
		if predicate(key, value) {
			filteredMap[key] = value
		}
	}
	return filteredMap
}

// MapValues creates a new map changing the values using the passing predicate
func (m GenericMap[K, V]) MapValues(predicate func(K, V) V) GenericMap[K, V] {
	mapped := New[K, V]()
	for key, value := range m {
		mapped.Add(key, predicate(key, value))
	}
	return mapped
}

// AddAll add all elements contained in the passed map
func (m GenericMap[K, V]) AddAll(origin map[K]V) GenericMap[K, V] {
	for k, v := range origin {
		m.Add(k, v)
	}
	return m
}
