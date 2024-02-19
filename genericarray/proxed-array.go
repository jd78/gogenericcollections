package genericarray

import (
	"github.com/jd78/gogenericcollections/composition/arraycomposition"
)

type ProxedArray[K comparable] struct {
	genericArray *GenericArray[K]
	composition  *arraycomposition.ArrayComposition[K]
}

func NewProxedArray[K comparable](g *GenericArray[K], ac *arraycomposition.ArrayComposition[K]) *ProxedArray[K] {
	return &ProxedArray[K]{g, ac}
}

func (pa *ProxedArray[K]) Filter(predicate arraycomposition.Filter[K]) *ProxedArray[K] {
	pa.composition.AddFunction(predicate)
	return pa
}

func (pa *ProxedArray[K]) MapValues(predicate arraycomposition.MapValues[K]) *ProxedArray[K] {
	pa.composition.AddFunction(predicate)
	return pa
}

func (pa *ProxedArray[K]) ToArray() GenericArray[K] {
	a := New[K]()
	for _, value := range *pa.genericArray {
		shouldAdd := true
		newVal := value
		for _, p := range pa.composition.GetPredicates() {
			switch predicate := p.(type) {
			case arraycomposition.Filter[K]:
				if !predicate.Exec(newVal) {
					shouldAdd = false
					break
				}
			case arraycomposition.MapValues[K]:
				newVal = predicate.Exec(newVal)
			}
		}
		if shouldAdd {
			a.Add(newVal)
		}
	}
	return *a
}
