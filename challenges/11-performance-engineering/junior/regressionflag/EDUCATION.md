# Deciding When A Number Is Bad News

## Intuition

Three outcomes, two thresholds, one symmetric band around zero. Everything inside the band is the machine, not the code.

## Approach

1. Clamp a negative tolerance to zero.
2. Compare against `+tolerance` and `-tolerance`, treating the boundaries as noise.
3. `Failing` classifies each change and reports the first regression.

## Solution

```go
func Classify(percent, tolerance float64) string {
	if tolerance < 0 {
		tolerance = 0
	}
	switch {
	case percent > tolerance:
		return "regression"
	case percent < -tolerance:
		return "improvement"
	default:
		return "noise"
	}
}

func Failing(percents []float64, tolerance float64) bool {
	for _, p := range percents {
		if Classify(p, tolerance) == "regression" {
			return true
		}
	}
	return false
}
```

## Walkthrough

With a zero tolerance the band collapses to the single point `0`, so `+0.1%` is a regression and an unchanged benchmark is still noise — which is why the boundary must be inclusive.

## Pitfalls

- `>=` at the boundary, which flags a change sitting exactly on the documented tolerance.
- Averaging the changes and testing the mean, which lets one bad benchmark hide behind nine good ones.
- Setting the tolerance from a single noisy run rather than from the observed spread.
