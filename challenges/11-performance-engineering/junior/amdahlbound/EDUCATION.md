# The Ceiling On Every Optimisation

## Intuition

Split the runtime into the part you touched and the part you did not. The second part does not move, so it becomes the whole cost of the program in the limit.

## Approach

1. Validate `p` and `s`, returning `1` for anything meaningless.
2. Apply the formula.
3. `Ceiling` special-cases `p >= 1` and otherwise returns `1/(1-p)`.

## Solution

```go
func MaxSpeedup(p, s float64) float64 {
	if p < 0 || p > 1 || s < 1 {
		return 1
	}
	return 1 / ((1 - p) + p/s)
}

func Ceiling(p float64) float64 {
	if p >= 1 {
		return math.Inf(1)
	}
	if p < 0 {
		return 1
	}
	return 1 / (1 - p)
}
```

## Walkthrough

`MaxSpeedup(0.02, 10)` is `1/(0.98 + 0.002)` — about 1.018. Ten times faster on a fiftieth of the work buys under two percent.

## Pitfalls

- Reporting the local speedup `s` as the program's speedup.
- Dividing by zero when `p` is exactly 1 without the guard.
- Estimating `p` by intuition instead of from a profile — it is almost always smaller than it feels.
