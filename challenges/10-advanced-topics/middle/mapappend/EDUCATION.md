# The Append That Never Reached The Map

## Intuition

Reading `m[key]` copies the slice header out of the map. `append` may reallocate, and even when it does not, the new length lives only in the copy — so unless you store it back, the map keeps the old header.

## Approach

1. Guard the nil map.
2. `m[key] = append(m[key], v)`.

## Solution

```go
// Add appends v to the slice stored under key, creating the entry when it
// is missing.
//
// A map value is not addressable: appending to m[key] produces a new slice
// header that has to be stored back.
//
// Examples:
//
// 	m := map[string][]int{}; Add(m, "a", 1) => m["a"] is [1]
func Add(m map[string][]int, key string, v int) {
	if m == nil {
		return
	}
	m[key] = append(m[key], v)
}
```

## Walkthrough

For a missing key, `m[key]` is a nil slice; `append` allocates and returns a one-element slice, which the assignment stores. Without the assignment the new slice is discarded and the key never appears.

## Pitfalls

- Checking `if _, ok := m[key]; !ok` and pre-creating the slice — harmless, and it does not fix the missing assignment.
- Assuming a map of pointers would behave the same; there the value is a pointer you can write through.
