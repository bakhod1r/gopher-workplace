# The Two Percentage Columns

## Intuition

The first column is a ratio per row; the second is that ratio accumulated. Both come from the same total, computed once before the listing is built.

## Approach

1. Sum the positive values.
2. Sort the rows, accumulate an exact running total, and round both percentages into each row.
3. `CoveringCount` walks the sorted rows until the running share reaches the target.

## Solution

```go
func Top(flat map[string]int64) []Row {
	var total int64
	for _, v := range flat {
		if v > 0 {
			total += v
		}
	}
	out := make([]Row, 0, len(flat))
	if total == 0 {
		return out
	}
	for fn, v := range flat {
		if v > 0 {
			out = append(out, Row{Func: fn, Flat: v})
		}
	}
	slices.SortFunc(out, func(a, b Row) int {
		if c := cmp.Compare(b.Flat, a.Flat); c != 0 {
			return c
		}
		return cmp.Compare(a.Func, b.Func)
	})
	var running int64
	for i := range out {
		running += out[i].Flat
		out[i].Pct = round2(float64(out[i].Flat) / float64(total) * 100)
		out[i].CumPct = round2(float64(running) / float64(total) * 100)
	}
	return out
}

func round2(v float64) float64 { return math.Round(v*100) / 100 }

func CoveringCount(flat map[string]int64, pct float64) int {
	if pct <= 0 {
		return 0
	}
	rows := Top(flat)
	for i, r := range rows {
		if r.CumPct >= pct {
			return i + 1
		}
	}
	return len(rows)
}
```

## Walkthrough

Three equal rows each round to 33.33, which sums to 99.99 — but because `CumPct` is computed from the exact running total rather than from the rounded columns, the final row still reads 100.

## Pitfalls

- Accumulating the rounded percentages, so the last row lands at 99.99.
- Computing `CumPct` before sorting, which makes the column meaningless.
- Reading a covering count from an unnormalised profile taken over a different duration.
