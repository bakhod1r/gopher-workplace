# Ordering A Report Without Wrecking The Data

## Intuition

A slice is a view onto an array the caller also holds. Sorting it rewrites their data, so a function that promises a sorted *copy* must actually copy.

## Approach

1. Clone the input.
2. Sort the clone with a three-key comparison.

## Solution

```go
func SortByCum(entries []Entry) []Entry {
	out := slices.Clone(entries)
	if out == nil {
		out = []Entry{}
	}
	slices.SortFunc(out, func(a, b Entry) int {
		if c := cmp.Compare(b.Cum, a.Cum); c != 0 {
			return c
		}
		if c := cmp.Compare(b.Flat, a.Flat); c != 0 {
			return c
		}
		return cmp.Compare(a.Func, b.Func)
	})
	return out
}
```

## Walkthrough

`slices.Clone(nil)` returns nil, so the empty case still needs one line to satisfy the non-nil contract.

## Pitfalls

- Sorting `entries` directly and returning it — the caller's order is gone.
- Comparing only `Cum`, leaving equal rows in input order rather than the documented one.
- Returning the clone before sorting it.
