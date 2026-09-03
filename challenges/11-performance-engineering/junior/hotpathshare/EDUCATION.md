# How Much Of The Profile Is This?

## Intuition

The share is a ratio between two sums drawn from the same map: the selected functions over everything.

## Approach

1. Build a set from the names.
2. Walk the map once, adding to the total always and to the selected sum when the key is in the set.
3. Guard the zero total before dividing.

## Solution

```go
func Share(flat map[string]int64, names []string) float64 {
	want := make(map[string]bool, len(names))
	for _, n := range names {
		want[n] = true
	}
	var total, sel int64
	for fn, v := range flat {
		total += v
		if want[fn] {
			sel += v
		}
	}
	if total <= 0 {
		return 0
	}
	return float64(sel) / float64(total)
}
```

## Walkthrough

Walking the map (not the names) is what makes duplicates harmless: each key is visited exactly once regardless of how many times it was requested.

## Pitfalls

- Ranging the names and summing `flat[n]`, which double-counts repeats.
- Dividing before the zero-total guard.
- Reporting the share of an optimised function as the speedup you will get — the rest of the program does not get faster.
