# Detect Whether Append Reallocated

## Intuition

Whether `append` copied is invisible through the language's own API: the length and capacity change either way. The backing array's address is the one observable that distinguishes them.

## Approach

1. Handle the zero-capacity cases explicitly.
2. Otherwise compare `unsafe.SliceData(before)` with `unsafe.SliceData(after)`.

## Solution

```go
import "unsafe"

// Grew reports whether after occupies different storage from before —
// that is, whether the append that produced it had to reallocate.
//
// Examples:
//
// 	s := make([]int, 0, 1); Grew(s, append(s, 1)) => false
func Grew(before, after []int) bool {
	if cap(before) == 0 || cap(after) == 0 {
		return cap(before) != cap(after)
	}
	return unsafe.SliceData(before) != unsafe.SliceData(after)
}
```

## Walkthrough

Appending to a slice with spare capacity returns a header with the same data pointer, so the comparison is false. When the capacity runs out, `append` allocates and the pointers differ.

## Pitfalls

- Comparing capacities instead — growth changes the capacity, but so does a three-index reslice.
- Forgetting the nil case, where both data pointers are nil.
