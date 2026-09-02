# Compare

## Intuition

Two comparisons cover all three outcomes: anything that is neither less nor greater must be equal, which is exactly what `default` handles.

## Approach

1. Return `-1` when `a < b`.
2. Return `1` when `a > b`.
3. Otherwise return `0`.

## Solution

```go
func Compare[T cmp.Ordered](a, b T) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
```

## Walkthrough

`Compare("b", "a")` fails the first case, matches the second, and returns `1`.

## Pitfalls

- Returning `0` for the less-than case because the cases are ordered wrongly.
- Subtracting (`a - b`), which does not compile for strings and overflows for ints.
- Returning arbitrary magnitudes instead of exactly `-1`, `0`, `1`.
