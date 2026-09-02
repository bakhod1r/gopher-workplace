# Sliding Window

## Intuition

Copying each window matters more here than for chunking: adjacent windows overlap, so sub-slices would share elements and an append into one could corrupt the next.

## Approach

1. Return empty for a non-positive or oversized `n`.
2. Slide `i` while `i+n <= len(s)`, copying `s[i:i+n]` each time.

## Solution

```go
func Windows[T any](s []T, n int) [][]T {
	out := make([][]T, 0)
	if n <= 0 || n > len(s) {
		return out
	}
	for i := 0; i+n <= len(s); i++ {
		w := make([]T, n)
		copy(w, s[i:i+n])
		out = append(out, w)
	}
	return out
}
```

## Walkthrough

`Windows([]int{1,2,3}, 2)` emits `[1 2]` at `i = 0` and `[2 3]` at `i = 1`, then stops because `2+2 > 3`.

## Pitfalls

- Using `i < len(s)`, which produces short trailing windows.
- Appending `s[i:i+n]` directly, leaving overlapping aliases.
- Returning one empty window when `n > len(s)`.
