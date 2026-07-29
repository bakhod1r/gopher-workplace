# The single-branch if

## Intuition

An `if` without `else` handles the exceptional case and falls through to the common return.

## Approach

1. Negatives → negate.
2. Otherwise return as-is.

## Solution

```go
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
```

## Walkthrough

`Abs(-5)` returns `-(-5) = 5`.

## Pitfalls

- `math.MinInt` has no positive counterpart; ignore that edge for this junior task.
- No `else` is needed after a `return`.
