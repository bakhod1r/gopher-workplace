# What The Trace Actually Covers

## Intuition

Sorted by start, the intervals can be swept left to right: each one either overlaps what you are building or begins a new block.

## Approach

1. Copy, drop empty spans, sort by start.
2. Sweep: extend the current span's end with `max`, or emit it and start a new one.
3. `Covered` sums the merged durations.

## Solution

```go
func Merge(spans []Span) []Span {
	valid := make([]Span, 0, len(spans))
	for _, s := range spans {
		if s.End > s.Start {
			valid = append(valid, s)
		}
	}
	slices.SortFunc(valid, func(a, b Span) int { return cmp.Compare(a.Start, b.Start) })
	out := make([]Span, 0, len(valid))
	for _, s := range valid {
		if n := len(out); n > 0 && s.Start <= out[n-1].End {
			out[n-1].End = max(out[n-1].End, s.End)
			continue
		}
		out = append(out, s)
	}
	return out
}

func Covered(spans []Span) int64 {
	var total int64
	for _, s := range Merge(spans) {
		total += s.End - s.Start
	}
	return total
}
```

## Walkthrough

`max(out[n-1].End, s.End)` is what protects the nested case: `{0,100}` followed by `{10,20}` keeps the end at 100 instead of pulling it back to 20.

## Pitfalls

- Assigning `out[n-1].End = s.End`, which truncates the outer span whenever one is nested inside it.
- Using `s.Start < cur.End`, which leaves touching spans `{0,10}` and `{10,20}` unmerged.
- Sorting the caller's slice in place, so their trace comes back reordered.
