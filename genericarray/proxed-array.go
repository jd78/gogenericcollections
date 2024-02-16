package genericarray

import (
	arrayfilter "github.com/jd78/gogenericcollections/composition/filter/array-filter"
	arrayvalues "github.com/jd78/gogenericcollections/composition/transform/array-values"
)

type ProxedArray[K comparable] struct {
	genericArray GenericArray[K]
	filter       *arrayfilter.Filter[K]
	mapValues    *arrayvalues.MapValues[K]
}

func NewWithFilter[K comparable](g GenericArray[K], f *arrayfilter.Filter[K]) *ProxedArray[K] {
	return &ProxedArray[K]{g, f, arrayvalues.New[K]()}
}

func NewWithMapValues[K comparable](g GenericArray[K], m *arrayvalues.MapValues[K]) *ProxedArray[K] {
	return &ProxedArray[K]{g, arrayfilter.New[K](), m}
}

func (f *ProxedArray[K]) Filter(predicate func(K) bool) *ProxedArray[K] {
	f.filter.AddFilter(predicate)
	return f
}

func (f *ProxedArray[K]) MapValues(predicate func(K) K) *ProxedArray[K] {
	f.mapValues.Map(predicate)
	return f
}

func (f *ProxedArray[K]) ToArray() GenericArray[K] {
	composedFilters := f.filter.Compose()
	composedMapValues := f.mapValues.Compose()

	proxedArray := New[K]()
	for _, value := range f.genericArray {
		if composedFilters(value) {
			proxedArray.Add(composedMapValues(value))
		}
	}
	return *proxedArray
}
