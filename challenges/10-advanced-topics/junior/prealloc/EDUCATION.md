# Allocate The Slice Once

## Intuition

`append` cannot see the future: each time it runs out of capacity it allocates a larger array and copies everything over. When the final size is known, ask for it once.

## Approach

1. Allocate `make([]int, n)`.
2. Fill each index with `i * i`.
3. Return the slice.

## Solution

```go
// Squares returns the squares 0..n-1 in order.
//
// The result must be built with a single allocation: give the slice its
// final length up front instead of growing it element by element.
//
// Examples:
//
// 	Squares(3) => []int{0, 1, 4}
func Squares(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = i * i
	}
	return out
}
```

## Walkthrough

For n = 64, appending from a nil slice reallocates about seven times (1, 2, 4, 8, 16, 32, 64) and copies 63 elements in total. `make([]int, 64)` allocates once and copies nothing.

## Pitfalls

- `make([]int, n)` versus `make([]int, 0, n)` — the first is already filled with zeros and is indexed, the second must be appended to.
- Using `append` on a slice made with a length appends *after* the zeros.
