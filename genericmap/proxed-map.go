package genericmap

import (
	"github.com/jd78/gogenericcollections/composition/mapcomposition"
)

type ProxedMap[K comparable, V any] struct {
	genericMap  GenericMap[K, V]
	composition *mapcomposition.MapComposition[K, V]
}

func newProxedMap[K comparable, V any](g GenericMap[K, V], ac *mapcomposition.MapComposition[K, V]) *ProxedMap[K, V] {
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

func (pm *ProxedMap[K, V]) Limit(limit int) *ProxedMap[K, V] {
	pm.composition.AddLimit(limit)
	return pm
}

func (pm *ProxedMap[K, V]) ToMap() GenericMap[K, V] {
	added := 0
	a := New[K, V]()
	for key, value := range pm.genericMap {
		shouldAdd := true
		newVal := value
		for _, p := range pm.composition.GetPredicates() {
			switch predicate := p.(type) {
			case mapcomposition.Filter[K, V]:
				if !predicate.Exec(key, newVal) {
					shouldAdd = false
					break
				}
			case mapcomposition.MapValues[K, V]:
				newVal = predicate.Exec(key, newVal)
			}
		}
		if shouldAdd {
			added++
			a.Add(key, newVal)
		}
		if pm.composition.GetLimit() > 0 && added >= pm.composition.GetLimit() {
			return a
		}
	}
	return a
}
