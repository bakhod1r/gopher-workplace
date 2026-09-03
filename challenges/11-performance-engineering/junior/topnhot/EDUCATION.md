# The Top Listing

## Intuition

Sorting turns an unordered map into a report. The only subtlety is that ties need a second key, or the randomised map order leaks into the output.

## Approach

1. Copy the map into a slice of entries.
2. Sort by value descending, then by name ascending.
3. Clamp `n` and reslice.

## Solution

```go
func TopN(flat map[string]int64, n int) []Entry {
	entries := make([]Entry, 0, len(flat))
	for fn, v := range flat {
		entries = append(entries, Entry{fn, v})
	}
	slices.SortFunc(entries, func(a, b Entry) int {
		if c := cmp.Compare(b.Value, a.Value); c != 0 {
			return c
		}
		return cmp.Compare(a.Func, b.Func)
	})
	if n < 0 {
		n = 0
	}
	return entries[:min(n, len(entries))]
}
```

## Walkthrough

`cmp.Compare(b.Value, a.Value)` with the arguments swapped is descending order; the fallthrough to `cmp.Compare(a.Func, b.Func)` keeps equal values in name order.

## Pitfalls

- Sorting on value alone, which leaves ties in randomised order.
- `entries[:n]` without clamping, which panics when the profile is short.
- Returning a nil slice for `n <= 0`.
