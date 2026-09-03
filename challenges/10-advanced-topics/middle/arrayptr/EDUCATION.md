# A Pointer To An Array Is Not A Slice

## Intuition

Go lets a pointer to an array be indexed and ranged as if it were the array. That gives you the cheapness of a pointer with the compile-time length check a slice cannot provide.

## Approach

1. Return 0 for a nil pointer.
2. Range the pointer and accumulate.

## Solution

```go
// Sum totals the array a points at.
//
// A pointer to an array carries the length in its type, so it can be
// indexed and ranged directly, with no header and no allocation.
//
// Examples:
//
// 	Sum(&[8]int{1, 2}) => 3
func Sum(a *[8]int) int {
	if a == nil {
		return 0
	}
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
```

## Walkthrough

Passing `&a` moves one word. `range a` iterates the eight elements in the caller's storage — no copy, no header, no allocation.

## Pitfalls

- `(*a)[i]` works but is noise; the pointer indexes directly.
- Forgetting the nil check — ranging a nil array pointer panics.
