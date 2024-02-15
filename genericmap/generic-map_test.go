package genericmap

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGenericMap_HasKey(t *testing.T) {
    m := NewGenericMap[string, int]()
    m.Add("one", 1)
    assert.True(t, m.HasKey("one"))
}

func TestGenericMap_Add(t *testing.T) {
    m := NewGenericMap[string, int]()

    m.Add("one", 1)

    expected := map[string]int{"one": 1}
    assert.Equal(t, expected, map[string]int(m))
}

func TestGenericMap_Get(t *testing.T) {
    m := NewGenericMap[string, string]()
    m.Add("key", "value")

    val := m.Get("key")
    assert.Equal(t, "value", val)
}

func TestGenericMap_Delete(t *testing.T) {
    m := NewGenericMap[string, int]()
    m.Add("one", 1)

    m.Delete("one")
    assert.False(t, m.HasKey("one"))
}

func TestGenericMap_Filter(t *testing.T) {
    m := NewGenericMap[string, int]()
    m.Add("one", 1)
    m.Add("two", 2)
    m.Add("three", 3)

    filtered := m.Filter(func(key string, value int) bool {
        return value%2 == 0
    })

    expected := map[string]int{"two": 2}
    assert.Equal(t, expected, map[string]int(filtered))
}

func TestGenericMap_MapValues(t *testing.T) {
    m := NewGenericMap[string, int]()
    m.Add("one", 1)
    m.Add("two", 2)
    m.Add("three", 3)

    mapped := m.MapValues(func(key string, value int) int {
        return value * 2
    })

    expected := map[string]int{"one": 2, "two": 4, "three": 6}
    assert.Equal(t, expected, map[string]int(mapped))
}

func TestGenericMap_AddAll(t *testing.T) {
    m1 := NewGenericMap[string, int]()
    m1.Add("one", 1)

    m2 := make(map[string]int)
    m2["two"] = 2
    m2["three"] = 3

    m1.AddAll(m2)

    expected := map[string]int{"one": 1, "two": 2, "three": 3}
    assert.Equal(t, expected, map[string]int(m1))
}

package extensions

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