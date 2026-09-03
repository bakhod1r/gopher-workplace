# Merging Histograms Across Instances

## Intuition

Two histograms with the same edges are two counts of the same partition, so they add elementwise. Different edges describe different partitions, and adding them is meaningless.

## Approach

1. `Valid` checks the length relationship and strict ascent.
2. `Merge` validates both, compares bounds, and sums into new slices.

## Solution

```go
func Valid(h Hist) bool {
	if len(h.Counts) != len(h.Bounds)+1 {
		return false
	}
	for i := 1; i < len(h.Bounds); i++ {
		if h.Bounds[i] <= h.Bounds[i-1] {
			return false
		}
	}
	return true
}

func Merge(a, b Hist) (Hist, bool) {
	if !Valid(a) || !Valid(b) || !slices.Equal(a.Bounds, b.Bounds) {
		return Hist{}, false
	}
	out := Hist{
		Bounds: slices.Clone(a.Bounds),
		Counts: make([]int64, len(a.Counts)),
	}
	for i := range a.Counts {
		out.Counts[i] = a.Counts[i] + b.Counts[i]
	}
	return out, true
}
```

## Walkthrough

Cloning the bounds matters as much as summing the counts: returning `a.Bounds` directly would hand the caller a slice that later writes to the merged histogram could edit under them.

## Pitfalls

- Comparing bounds with `==` on the slices, which does not compile, or by length only, which accepts different edges.
- Summing into `a.Counts`, mutating an input.
- Merging histograms whose bounds "look close enough" — the counts become fiction.
