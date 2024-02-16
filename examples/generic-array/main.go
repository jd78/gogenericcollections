package main

import (
	"fmt"
	"strconv"

	"github.com/jd78/gogenericcollections/genericarray"
)

func main() {
	a := genericarray.New[int]()

	a.Add(1)
	a.Add(2)

	for k := range a {
		println(k)
	}

	aplus1 := a.MapValues(func(i int) int {
		return i + 1
	})

	for k := range aplus1.ToArray() {
		println(k)
	}

	// transforms the array of int to int64
	genericarray.MapArray[int, int64](a, func(i int) int64 {
		return int64(i)
	})

	m := genericarray.ArrayToMap[int, int, string](a, func(i int) (int, string) {
		return i, strconv.Itoa(i)
	})

	for k, v := range m {
		fmt.Printf("k=%d, v=%s\n", k, v)
	}
}
