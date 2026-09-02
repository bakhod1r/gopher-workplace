# Repeat From Stdlib

## Intuition

Reading the documented panics before delegating is the habit here; the wrapper exists to convert them into a total function.

## Approach

1. Return an empty slice for `n <= 0`.
2. Otherwise return `slices.Repeat(s, n)`.

## Solution

```go
func Tile[T any](s []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	return slices.Repeat(s, n)
}
```

## Walkthrough

`Tile([]int{1}, -1)` never reaches `slices.Repeat`, so the panic cannot happen.

## Pitfalls

- Passing a negative count through and panicking.
- Returning nil for `n == 0`.
- Repeating by hand when the stdlib helper exists.
