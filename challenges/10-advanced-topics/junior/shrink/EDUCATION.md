# Give Back The Capacity You Stopped Using

## Intuition

The collector frees allocations, not the unused part of one. Once a slice is a small window on a big array, the only way to release the rest is to copy the window somewhere its own size.

## Approach

1. Compare `cap(s)` with `2*len(s)`; if it is not bigger, return `s`.
2. Otherwise allocate `make([]int, len(s))`, copy, return the copy.

## Solution

```go
// Shrink returns a copy of s sized exactly to its length when s is
// holding on to far more capacity than it uses, and returns s unchanged
// otherwise.
//
// "Far more" means the capacity is more than twice the length.
//
// Examples:
//
// 	Shrink(make([]int, 2, 64)) => a slice of length 2 and capacity 2
func Shrink(s []int) []int {
	if cap(s) <= 2*len(s) {
		return s
	}
	out := make([]int, len(s))
	copy(out, s)
	return out
}
```

## Walkthrough

A slice of len 2, cap 64 wastes 62 slots. Copying two ints into a two-int array lets the 64-int array be collected. A slice of len 8, cap 10 fails the test and is returned as is.

## Pitfalls

- Shrinking unconditionally — every call then allocates.
- `s[:len(s):len(s)]` caps the capacity but keeps pointing at the same big array.
