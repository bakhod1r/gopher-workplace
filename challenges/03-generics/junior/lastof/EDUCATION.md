# Last Element

## Intuition

The generic part is trivial; the trap is the same as in any Go slice code — `len(s)-1` is only a valid index when the slice is non-empty.

## Approach

1. Return `zero, false` when the slice is empty.
2. Otherwise return `s[len(s)-1], true`.

## Solution

```go
func Last[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[len(s)-1], true
}
```

## Walkthrough

`Last([]int{3, 1, 4})` computes index `2` and returns `4`. `Last([]int{})` would compute index `-1`, so the guard runs first.

## Pitfalls

- Using `s[len(s)]` — off by one, always out of range.
- Skipping the empty guard, producing an index-out-of-range panic.
- Assuming `T` can be compared to `0` to detect emptiness.
