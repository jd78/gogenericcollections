package genericmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericMap_HasKey(t *testing.T) {
	m := New[string, int]()
	m.Add("one", 1)
	assert.True(t, m.HasKey("one"))
}

func TestGenericMap_Add(t *testing.T) {
	m := New[string, int]()

	m.Add("one", 1)

	expected := map[string]int{"one": 1}
	assert.Equal(t, expected, map[string]int(m))
}

func TestGenericMap_Get(t *testing.T) {
	m := New[string, string]()
	m.Add("key", "value")

	val := m.Get("key")
	assert.Equal(t, "value", val)
}

func TestGenericMap_Delete(t *testing.T) {
	m := New[string, int]()
	m.Add("one", 1)

	m.Delete("one")
	assert.False(t, m.HasKey("one"))
}

func TestGenericMap_Filter(t *testing.T) {
	m := New[string, int]()
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	filtered := m.Filter(func(key string, value int) bool {
		return value >= 1
	}).MapValues(func(key string, value int) int {
		return value * 2
	}).Filter(func(key string, value int) bool {
		return value < 3
	})

	expected := map[string]int{"one": 2}
	assert.Equal(t, expected, map[string]int(filtered.ToMap()))
}

func TestGenericMap_MapValues(t *testing.T) {
	m := New[string, int]()
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	mapped := m.MapValues(func(key string, value int) int {
		return value * 2
	})

	expected := map[string]int{"one": 2, "two": 4, "three": 6}
	assert.Equal(t, expected, map[string]int(mapped.ToMap()))
}

func TestGenericMap_AddAll(t *testing.T) {
	m1 := New[string, int]()
	m1.Add("one", 1)

	m2 := make(map[string]int)
	m2["two"] = 2
	m2["three"] = 3

	m1.AddAll(m2)

	expected := map[string]int{"one": 1, "two": 2, "three": 3}
	assert.Equal(t, expected, map[string]int(m1))
}

func TestGenericMap_Composition(t *testing.T) {
	m := New[string, int]().
		Add("one", 1).
		Add("two", 2).
		Add("three", 3).
		Filter(func(k string, v int) bool {
			if v < 3 {
				return true
			}
			return false
		}).MapValues(func(key string, value int) int {
		return value * 2
	})

	expected := map[string]int{"one": 2, "two": 4}
	assert.Equal(t, expected, map[string]int(m.ToMap()))
}
