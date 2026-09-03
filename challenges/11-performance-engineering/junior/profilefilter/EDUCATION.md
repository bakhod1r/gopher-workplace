# Hiding The Noise

## Intuition

Every row is either above the cut or below it. Compute the cut once, then keep the rows that clear it.

## Approach

1. Guard a non-positive total.
2. Compute the absolute threshold from the percentage.
3. Append qualifying rows to a fresh slice.

## Solution

```go
func Filter(entries []Entry, total int64, minPct float64) []Entry {
	out := make([]Entry, 0, len(entries))
	if total <= 0 {
		return out
	}
	cut := float64(total) * minPct / 100
	for _, e := range entries {
		if float64(e.Value) >= cut {
			out = append(out, e)
		}
	}
	return out
}
```

## Walkthrough

Computing `cut` once outside the loop keeps the comparison a single float compare per row, and keeps the boundary consistent for every entry.

## Pitfalls

- A strict `>`, which drops a node sitting exactly on the threshold.
- Filtering in place with `entries[:0]`, which overwrites the caller's slice.
- Comparing percentages per row with repeated division, which invites rounding drift.
