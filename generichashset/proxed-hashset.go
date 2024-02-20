package generichashset

import (
	"github.com/jd78/gogenericcollections/composition/hashsetcomposition"
)

type ProxedHashset[K comparable] struct {
	genericHashset GenericHashset[K]
	composition    *hashsetcomposition.HashsetComposition[K]
}

func newProxedHashset[K comparable](g GenericHashset[K], ac *hashsetcomposition.HashsetComposition[K]) *ProxedHashset[K] {
	return &ProxedHashset[K]{g, ac}
}

func (pm *ProxedHashset[K]) Filter(predicate hashsetcomposition.Filter[K]) *ProxedHashset[K] {
	pm.composition.AddFunction(predicate)
	return pm
}

func (pm *ProxedHashset[K]) ToHashset() GenericHashset[K] {
	a := New[K]()
	for key := range pm.genericHashset {
		shouldAdd := true
		newKey := key
		for _, p := range pm.composition.GetPredicates() {
			switch predicate := p.(type) {
			case hashsetcomposition.Filter[K]:
				if !predicate.Exec(newKey) {
					shouldAdd = false
					break
				}
			}
		}
		if shouldAdd {
			a.Add(newKey)
		}
	}
	return a
}
