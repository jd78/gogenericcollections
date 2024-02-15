package genericmap

type GenericMap[K comparable, V any] map[K]V

func New[K comparable, V any]() GenericMap[K, V] {
    return make(GenericMap[K, V])
}

func (m GenericMap[K, V]) HasKey(key K) bool {
    _, exists := m[key]
    return exists
}

func (m GenericMap[K, V]) Add(key K, value V) {
    m[key] = value
}

func (m GenericMap[K, V]) Get(key K) V {
    return m[key]
}

func (m GenericMap[K, V]) Delete(key K) {
    if !m.HasKey(key) {
        return
    }
    var zero V
    m[key] = zero
    delete(m, key)
}

func (m GenericMap[K, V]) Filter(predicate func(K, V) bool) GenericMap[K, V] {
    filteredMap := NewGenericMap[K, V]()
    for key, value := range m {
        if predicate(key, value) {
            filteredMap[key] = value
        }
    }
    return filteredMap
}

func (m GenericMap[K, V]) MapValues(predicate func(K, V) V) GenericMap[K, V] {
    mapped := NewGenericMap[K, V]()
    for key, value := range m {
        mapped.Add(key, predicate(key, value))
    }
    return mapped
}

func (m GenericMap[K, V]) AddAll(origin map[K]V) GenericMap[K, V] {
    for k, v := range origin {
        m.Add(k, v)
    }
    return m
}
