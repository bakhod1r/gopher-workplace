# Deep-copying maps with reference values

## The idea

A map copy that assigns each value copies **references** for reference-typed
values (slices, maps, pointers). To fully separate, clone the values too:

```go
for k, v := range m { out[k] = append([]int{}, v...) }
```

## Why it matters

`maps.Clone` and manual map copies are shallow. Mutating a slice in the "clone"
then reaches back into the original — the same shallow-vs-deep trap as struct
slice fields, now through map values.

## Watch out

- Every reference-typed value needs its own deep copy.
- Nested maps of maps need recursion.
- Value types (int, string, array) copy fully and are safe.
