package genericarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericArray_Add(t *testing.T) {
	array := NewGenericArray[int]()

	array.Add(1)

	expected := []int{1}
	assert.ElementsMatch(t, expected, array)
}

func TestGenericArray_AddAll(t *testing.T) {
	origin := []int{1, 2, 3}
	array := NewGenericArrayFrom(origin)

	array.AddAll([]int{4, 5})

	expected := []int{1, 2, 3, 4, 5}
	assert.ElementsMatch(t, expected, []int(array))
}

func TestGenericArray_MapValues(t *testing.T) {
	array := NewGenericArray[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	mapped := array.MapValues(func(val int) int {
		return val * 2
	})

	expected := []int{2, 4, 6}
	assert.ElementsMatch(t, expected, []int(mapped))
}

func TestMapArray(t *testing.T) {
	array := NewGenericArray[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	mapped := MapArray(array, func(val int) string {
		return fmt.Sprintf("%d+", val)
	})

	expected := []string{"1+", "2+", "3+"}
	assert.ElementsMatch(t, expected, []string(mapped))
}

func TestArrayToMap(t *testing.T) {
	array := NewGenericArray[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	m := ArrayToMap(array, func(val int) (int, string) {
		return val, fmt.Sprintf("%d", val)
	})

	expected := map[int]string{1: "1", 2: "2", 3: "3"}
	assert.Equal(t, expected, map[int]string(m))
}
