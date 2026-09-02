# Absolute Value

## Intuition

If `~uint` were in the set, `v < 0` would be always false and `-v` would wrap around. Keeping the constraint signed makes the implementation correct for every instantiation.

## Approach

1. Return `-v` when `v < 0`.
2. Otherwise return `v`.

## Solution

```go
func Abs[T Signed](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
```

## Walkthrough

`Abs(-2.5)` instantiates `T = float64`, takes the negative branch, and returns `2.5`.

## Pitfalls

- Adding `~uint` to the constraint, which makes the function silently wrong.
- Comparing against a literal that does not fit every type in the set.
- Returning `float64` and breaking integer callers.
