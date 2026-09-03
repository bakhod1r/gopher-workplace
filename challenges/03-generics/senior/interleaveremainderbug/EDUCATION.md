# Interleave That Drops The Tail

## Intuition

The loop stops at the shorter length and nothing collects the surplus, so up to `abs(len(a)-len(b))` elements are lost.

## Approach

1. Interleave up to the shorter length.
2. Append `a[n:]` and `b[n:]` — one of them is empty.

## Solution

```go
func Interleave[T any](a, b []T) []T {
	out := make([]T, 0, len(a)+len(b))
	n := min(len(a), len(b))
	for i := 0; i < n; i++ {
		out = append(out, a[i], b[i])
	}
	out = append(out, a[n:]...)
	out = append(out, b[n:]...)
	return out
}
```

## Walkthrough

`Interleave([1,2,3],[9])` emits `[1 9]` and drops `2` and `3`.

## Pitfalls

- Guessing which slice is longer instead of appending both tails.
- Indexing past the end when the loop bound uses `max`.
