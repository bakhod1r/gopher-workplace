# Making Two Profiles Comparable

## Intuition

Raw sample counts are in units of "however long I profiled". Dividing by the duration gives work per second; dividing by the total gives the profile's shape.

## Approach

1. `Rate` guards the duration and divides each value.
2. `Fractions` sums the positive values first, then divides in a second pass.

## Solution

```go
func Rate(flat map[string]int64, seconds float64) map[string]float64 {
	out := make(map[string]float64, len(flat))
	if seconds <= 0 {
		return out
	}
	for fn, v := range flat {
		out[fn] = float64(v) / seconds
	}
	return out
}

func Fractions(flat map[string]int64) map[string]float64 {
	var total int64
	for _, v := range flat {
		if v > 0 {
			total += v
		}
	}
	out := make(map[string]float64, len(flat))
	if total == 0 {
		return out
	}
	for fn, v := range flat {
		if v <= 0 {
			continue
		}
		out[fn] = float64(v) / float64(total)
	}
	return out
}
```

## Walkthrough

Skipping non-positive values in *both* passes is what keeps the fractions summing to exactly 1; counting them in the total but not in the output would leave the shares short.

## Pitfalls

- Summing the total including negative or zero entries, so the fractions no longer sum to 1.
- Comparing two profiles' raw values when one ran twice as long.
- Normalising to fractions and then claiming absolute improvement — shares can shift with no change in total work.
