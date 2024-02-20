package main

import (
	"fmt"

	"github.com/jd78/gogenericcollections/generichashset"
)

func main() {
	m := generichashset.New[int]()
	m.Add(1)
	m.Add(2)
	m.Add(3)

	fmt.Printf("has key 1 %v\n", m.HasKey(1))
	fmt.Printf("has non existing 5 %v\n", m.HasKey(5))
	m.Delete(1)

	fmt.Printf("has key 1 %v\n", m.HasKey(1))

	f := m.Filter(func(i int) bool { return i == 2 })

	for k := range f.ToHashset() {
		fmt.Printf("filtered k=%d\n", k)
	}
}
