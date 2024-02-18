package genericmap

import (
	"github.com/jd78/gogenericcollections/composition"
	"github.com/jd78/gogenericcollections/composition/mapcomposition"
)

type ProxedMap[K comparable, V any] struct {
	genericMap  GenericMap[K, V]
	composition *mapcomposition.MapComposition[K, V]
}

func NewProxedMap[K comparable, V any](g GenericMap[K, V], ac *mapcomposition.MapComposition[K, V]) *ProxedMap[K, V] {
	return &ProxedMap[K, V]{g, ac}
}

func (pm *ProxedMap[K, V]) Filter(predicate mapcomposition.Filter[K, V]) *ProxedMap[K, V] {
	pm.composition.AddFunction(predicate)
	return pm
}

func (pm *ProxedMap[K, V]) MapValues(predicate mapcomposition.MapValues[K, V]) *ProxedMap[K, V] {
	pm.composition.AddFunction(predicate)
	return pm
}

func (pm *ProxedMap[K, V]) ToMap() GenericMap[K, V] {
	// composedFilters := f.filter.Compose()
	// composedMapValues := f.mapValues.Compose()

	// proxedMap := New[K, V]()
	// for key, value := range f.genericMap {
	// 	if composedFilters(key, value) {
	// 		proxedMap[key] = composedMapValues(key, value)
	// 	}
	// }
	// return proxedMap

	a := New[K, V]()
	for key, value := range pm.genericMap {
		shouldAdd := true
		newVal := value
		for _, p := range pm.composition.GetPredicates() {
			switch p.GetType() {
			case composition.Filter:
				if !p.(mapcomposition.Filter[K, V]).Exec(key, newVal) {
					shouldAdd = false
					break
				}
			case composition.MapValues:
				newVal = p.(mapcomposition.MapValues[K, V]).Exec(key, newVal)
			}
		}
		if shouldAdd {
			a.Add(key, newVal)
		}
	}
	return a
}
