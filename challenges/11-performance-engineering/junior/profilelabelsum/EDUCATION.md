# Slicing A Profile By Label

## Intuition

Grouping by a label is the same aggregation as grouping by function name; only the key changes.

## Approach

1. Allocate the result map.
2. Skip non-positive values.
3. Look up the label — a missing key gives `""` for free — and accumulate.

## Solution

```go
func SumByLabel(samples []Sample, key string) map[string]int64 {
	out := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 {
			continue
		}
		out[s.Labels[key]] += s.Value
	}
	return out
}
```

## Walkthrough

`s.Labels[key]` on a nil `Labels` map is legal and returns `""`, so the nil case and the missing-key case collapse into the same branch-free line.

## Pitfalls

- Guarding with `if s.Labels == nil { continue }`, which silently loses unlabelled work.
- Using the comma-ok form and then skipping the sample instead of bucketing it under `""`.
- Writing to a nil result map, which panics.
