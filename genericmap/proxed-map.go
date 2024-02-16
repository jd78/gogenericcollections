package genericmap

import (
	mapfilter "github.com/jd78/gogenericcollections/composition/filter/map-filter"
	mapvalues "github.com/jd78/gogenericcollections/composition/transform/map-values"
)

type ProxedMap[K comparable, V any] struct {
	genericMap GenericMap[K, V]
	filter     *mapfilter.Filter[K, V]
	mapValues  *mapvalues.MapValues[K, V]
}

func NewWithFilter[K comparable, V any](g GenericMap[K, V], f *mapfilter.Filter[K, V]) *ProxedMap[K, V] {
	return &ProxedMap[K, V]{g, f, mapvalues.New[K, V]()}
}

func NewWithMapValues[K comparable, V any](g GenericMap[K, V], m *mapvalues.MapValues[K, V]) *ProxedMap[K, V] {
	return &ProxedMap[K, V]{g, mapfilter.New[K, V](), m}
}

func (f *ProxedMap[K, V]) Filter(predicate func(K, V) bool) *ProxedMap[K, V] {
	f.filter.AddFilter(predicate)
	return f
}

func (f *ProxedMap[K, V]) MapValues(predicate func(K, V) V) *ProxedMap[K, V] {
	f.mapValues.Map(predicate)
	return f
}

func (f *ProxedMap[K, V]) ToMap() GenericMap[K, V] {
	composedFilters := f.filter.Compose()
	composedMapValues := f.mapValues.Compose()

	proxedMap := New[K, V]()
	for key, value := range f.genericMap {
		if composedFilters(key, value) {
			proxedMap[key] = composedMapValues(key, value)
		}
	}
	return proxedMap
}
