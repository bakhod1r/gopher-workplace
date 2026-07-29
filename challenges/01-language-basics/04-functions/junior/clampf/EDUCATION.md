# Bounding a value

## Intuition

Clamping composes a lower and upper limit into one operation used across UI, graphics, and numeric guards.

## Approach

1. Below `lo` → return `lo`.
2. Above `hi` → return `hi`.
3. Otherwise return `v`.

## Solution

```go
func Clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
```

## Walkthrough

`Clamp(99, 0, 10)`: 99 exceeds the ceiling, so 10 is returned.

## Pitfalls

- With `lo > hi` the range is empty; the task assumes `lo <= hi`.
- Use inclusive comparisons so the endpoints are reachable.
