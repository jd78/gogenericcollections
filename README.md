# gogenericcollections

## Generic Map example

```go

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
```

## Generic Array example

```go
    a := genericarray.New[int]()

	a.Add(1)
	a.Add(2)

	for k := range a {
		println(k)
	}

	aplus1 := a.MapValues(func(i int) int {
		return i + 1
	})

	for k := range aplus1 {
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
```