# Sampling Fractions And What They Cost You

## Intuition

One event in `n` was recorded, so each recorded event stands for `n`. If nothing was recorded because the profiler was off, no multiplication can recover the data.

## Approach

1. `Scale` rejects a non-positive fraction and multiplies otherwise.
2. `Estimate` returns early for a disabled profile, then aggregates scaled delays.
3. `Confidence` is a banded classification.

## Solution

```go
func Scale(count, delay int64, fraction int) (int64, int64, bool) {
	if fraction <= 0 {
		return 0, 0, false
	}
	if fraction == 1 {
		return count, delay, true
	}
	return count * int64(fraction), delay * int64(fraction), true
}

func Estimate(records []Contention, fraction int) map[string]int64 {
	out := make(map[string]int64)
	if fraction <= 0 {
		return out
	}
	for _, r := range records {
		if r.Count <= 0 || r.Delay < 0 {
			continue
		}
		_, delay, ok := Scale(r.Count, r.Delay, fraction)
		if !ok {
			continue
		}
		out[r.Site] += delay
	}
	return out
}

func Confidence(samples int64) string {
	switch {
	case samples < 10:
		return "low"
	case samples < 100:
		return "medium"
	default:
		return "high"
	}
}
```

## Walkthrough

Returning `ok` rather than a zero from `Scale` is what separates "the profile says there was no contention" from "there was no profile" — two conclusions that look identical in the numbers and mean opposite things.

## Pitfalls

- Treating an empty profile as proof the locks are fine.
- Reporting a scaled delay from two samples as if it were measured.
- Running with fraction 1 in production, where the profiling overhead is itself a source of contention.
