# Reading A Mutex Profile

## Intuition

Group by call site, sum both columns, divide for the mean, sort by the column that matters. The interesting part is knowing which column that is.

## Approach

1. Aggregate valid records into a map keyed by site.
2. Flatten into rows, computing the mean.
3. Sort by delay descending, then site ascending.

## Solution

```go
func Rank(records []Record) []Site {
	type agg struct{ count, delay int64 }
	totals := make(map[string]agg, len(records))
	for _, r := range records {
		if r.Count <= 0 || r.Delay < 0 {
			continue
		}
		a := totals[r.Site]
		a.count += r.Count
		a.delay += r.Delay
		totals[r.Site] = a
	}
	out := make([]Site, 0, len(totals))
	for site, a := range totals {
		out = append(out, Site{
			Site:      site,
			Count:     a.count,
			Delay:     a.delay,
			MeanDelay: float64(a.delay) / float64(a.count),
		})
	}
	slices.SortFunc(out, func(x, y Site) int {
		if c := cmp.Compare(y.Delay, x.Delay); c != 0 {
			return c
		}
		return cmp.Compare(x.Site, y.Site)
	})
	return out
}

func Worst(records []Record) (Site, bool) {
	rows := Rank(records)
	if len(rows) == 0 {
		return Site{}, false
	}
	return rows[0], true
}
```

## Walkthrough

Dropping zero-count records before the aggregation is what makes `MeanDelay` safe: every surviving site has a positive count, so the division cannot produce a NaN.

## Pitfalls

- Ranking by blocking *count*, which promotes cheap frequent locks over genuinely slow ones.
- Dividing by a count that could be zero, producing NaN rows that sort unpredictably.
- Forgetting `runtime.SetMutexProfileFraction`, so the profile is empty and the absence of contention looks proven.
