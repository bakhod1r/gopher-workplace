# What Changed Between Two Profiles

## Intuition

The diff is a full outer join. Because a missing map key reads as `0`, the join costs nothing beyond collecting the key set.

## Approach

1. Collect the union of the two key sets.
2. For each key, read both sides and skip a zero delta.
3. Sort by `abs(delta)` descending, then name.

## Solution

```go
func Diff(base, candidate map[string]int64) []Change {
	names := make(map[string]bool, len(base)+len(candidate))
	for k := range base {
		names[k] = true
	}
	for k := range candidate {
		names[k] = true
	}
	out := make([]Change, 0, len(names))
	for name := range names {
		b, c := base[name], candidate[name]
		if c-b == 0 {
			continue
		}
		out = append(out, Change{Func: name, Base: b, New: c, Delta: c - b})
	}
	slices.SortFunc(out, func(x, y Change) int {
		if c := cmp.Compare(abs(y.Delta), abs(x.Delta)); c != 0 {
			return c
		}
		return cmp.Compare(x.Func, y.Func)
	})
	return out
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
```

## Walkthrough

`base[name]` for a function only present in the candidate returns `0`, which is precisely the semantics you want: it used to cost nothing and now it costs something.

## Pitfalls

- Ranging one map only, which silently drops every newly appeared function.
- Sorting by signed delta, which buries the biggest regression under the biggest win.
- Comparing raw values from profiles with different totals — normalise first, or the diff measures the workload, not the code.
