# gogenericcollections

WIP eventually more methods to come..

Streaming collections where you can chain functions like:

```go
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
```

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

for k, v := range f.ToMap() {
    fmt.Printf("filtered k=%d, v=%s\n", k, v)
}

mappedM := m.MapValues(func(i int, v string) string {
    return v + "edited"
})

for k, v := range mappedM.ToMap() {
    fmt.Printf("mapped k=%d, v=%s\n", k, v)
}
```

## Generic Array example

```go
a := genericarray.New[int]()
a.Add(1)
a.Add(2)

for k := range *a {
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
```

## Generic Hashset example

```go
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
```