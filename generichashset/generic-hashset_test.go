package generichashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericMap_HasKey(t *testing.T) {
	m := New[string]()
	m.Add("one")
	assert.True(t, m.HasKey("one"))
}

func TestGenericMap_Add(t *testing.T) {
	m := New[string]()

	m.Add("one")

	expected := map[string]struct{}{"one": struct{}{}}
	assert.Equal(t, expected, map[string]struct{}(m))
}

func TestGenericMap_Delete(t *testing.T) {
	m := New[string]()
	m.Add("one")

	m.Delete("one")
	assert.False(t, m.HasKey("one"))
}

func TestGenericMap_Filter(t *testing.T) {
	m := New[int]()
	m.Add(1)
	m.Add(2)
	m.Add(3)

	filtered := m.Filter(func(key int) bool {
		return key > 1
	})

	expected := map[int]struct{}{2: struct{}{}, 3: struct{}{}}
	assert.Equal(t, expected, map[int]struct{}(filtered.ToHashset()))
}

func TestGenericMap_AddAll(t *testing.T) {
	m1 := New[string]()
	m1.Add("one")

	m2 := make(map[string]struct{})
	m2["two"] = struct{}{}
	m2["three"] = struct{}{}

	m1.AddAll(m2)

	expected := map[string]struct{}{"one": struct{}{}, "two": struct{}{}, "three": struct{}{}}
	assert.Equal(t, expected, map[string]struct{}(m1))
}
