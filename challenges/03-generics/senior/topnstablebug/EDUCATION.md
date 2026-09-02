# Top-N That Sorts The Caller's Slice

## Intuition

Sorting `s` directly reorders the caller's backing array, and the returned prefix still points into it — two ownership violations from one line.

## Approach

1. Copy the input into a fresh slice.
2. Sort the copy descending with `slices.SortStableFunc`.
3. Clamp `n` into range and return the prefix of the copy.

## Solution

```go
func TopN[T any](s []T, score func(T) int, n int) []T {
	out := make([]T, len(s))
	copy(out, s)
	slices.SortStableFunc(out, func(a, b T) int {
		return score(b) - score(a)
	})
	if n > len(out) {
		n = len(out)
	}
	if n < 0 {
		n = 0
	}
	return out[:n]
}
```

## Walkthrough

After `TopN(rows, score, 2)` the caller's `rows` is in score order, so the arrival-order view is wrong.

## Pitfalls

- `score(b) - score(a)` overflows for extreme values; `cmp.Compare` is safer.
- Copying with `out := s`, which copies the header and not the data.
