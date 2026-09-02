# Averaging a Stream

## Intuition

A slice knows its length; a stream does not. So a streaming average keeps
two accumulators, the sum and the count, and combines them only at the end.
The empty window has no meaningful answer, which is what the `bool`
reports.

## Approach

1. Track `sum` (float64) and `n` (int).
2. `range` over `readings`, updating both.
3. If `n == 0`, return `0, false`.
4. Otherwise return `sum / float64(n), true`.

## Solution

```go
func AverageReading(readings <-chan float64) (float64, bool) {
	sum := 0.0
	n := 0
	for v := range readings {
		sum += v
		n++
	}
	if n == 0 {
		return 0, false
	}
	return sum / float64(n), true
}
```

## Walkthrough

For `1, 2, 3`: `sum` reaches 6 with `n == 3`, giving `6 / 3.0 == 2` and
`true`. For an empty window the loop never runs and the guard returns
`0, false`.

## Pitfalls

- `sum / n` does not compile — `n` is an `int` and needs an explicit `float64` conversion.
- Without the `n == 0` guard, floating-point division yields `NaN` rather than an error.
- `len(readings)` counts only what is buffered right now, not the whole window.
