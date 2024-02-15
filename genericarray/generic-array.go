package genericarray

type GenericArray[K comparable] []K

func NewGenericArray[K comparable]() GenericArray[K] {
    return make(GenericArray[K], 0)
}

func NewGenericArrayFrom[K comparable](origin []K) GenericArray[K] {
    a := make(GenericArray[K], 0)
    for _, v := range origin {
        a.Add(v)
    }
    return a
}

func (a *GenericArray[K]) Add(val K) {
    *a = append(*a, val)
}

func (a *GenericArray[K]) AddAll(origin []K) GenericArray[K] {
    for _, v := range origin {
        a.Add(v)
    }
    return *a
}

func (a GenericArray[K]) MapValues(predicate func(K) K) GenericArray[K] {
    mapped := NewGenericArray[K]()
    for _, v := range a {
        mapped.Add(predicate(v))
    }
    return mapped
}

func MapArray[K, V comparable](a GenericArray[K], predicate func(K) V) GenericArray[V] {
    mappedArray := NewGenericArray[V]()

    for _, item := range a {
        mappedItem := predicate(item)
        mappedArray.Add(mappedItem)
    }

    return mappedArray
}

func ArrayToMap[K, V, Z comparable](a GenericArray[K], predicate func(K) (V, Z)) GenericMap[V, Z] {
    m := NewGenericMap[V, Z]()

    for _, item := range a {
        first, second := predicate(item)
        m.Add(first, second)
    }

    return m
}