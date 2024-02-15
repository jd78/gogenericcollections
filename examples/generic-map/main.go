package main

import (
	"fmt"

	"github.com/jd78/gogenericcollections/genericmap"
)

func main() {
	m := genericmap.New[int, string]()
	m.Add(1, "test")
	m.Add(2, "test2")
	m.Add(3, "test3")

	fmt.Printf("get key 1 %v\n", m.Get(1))
	fmt.Printf("get non existing 5 %v\n", m.Get(5))

	fmt.Printf("has key 3 %v\n", m.HasKey(3))

	m.Delete(3)

	fmt.Printf("has key 3 %v\n", m.HasKey(3))

	f := m.Filter(func(i int, v string) bool { return i == 2 })

	for k, v := range f {
		fmt.Printf("filtered k=%d, v=%s\n", k, v)
	}

	mappedM := m.MapValues(func(i int, v string) string {
		return v + "edited"
	})

	for k, v := range mappedM {
		fmt.Printf("mapped k=%d, v=%s\n", k, v)
	}
}
