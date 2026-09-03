# Two Runs, One Table

## Intuition

The comparison is an inner join on the benchmark name, plus one percentage per matched pair.

## Approach

1. Range the base map, looking each name up in the candidate.
2. Skip unpaired names and unusable baselines.
3. Sort the collected rows by name.

## Solution

```go
func Compare(base, candidate map[string]float64) []Delta {
	out := make([]Delta, 0, len(base))
	for name, b := range base {
		if b <= 0 {
			continue
		}
		c, ok := candidate[name]
		if !ok {
			continue
		}
		out = append(out, Delta{Name: name, Base: b, New: c, Percent: (c - b) / b * 100})
	}
	slices.SortFunc(out, func(x, y Delta) int { return cmp.Compare(x.Name, y.Name) })
	return out
}
```

## Walkthrough

Ranging the base map and looking up the candidate gives the inner join in one pass; the sort afterwards is what makes the output identical on every run.

## Pitfalls

- Dividing by the candidate, which inverts the reported direction.
- Emitting rows for names present on only one side, comparing against an implicit zero.
- Skipping the sort and shipping a report that reorders itself every run.
