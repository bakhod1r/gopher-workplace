# Collapsing consecutive duplicates

## The idea

Emit an element only when it differs from the previously emitted one:

```go
out := []int{}
for i, v := range xs {
	if i == 0 || v != xs[i-1] { out = append(out, v) }
}
```

## Why it matters

On sorted data this removes all duplicates in O(n) with no map. Run-collapsing
also underlies run-length encoding and change detection.

## Watch out

- It only removes **consecutive** duplicates; unsorted input keeps distant repeats.
- `slices.Compact` (Go 1.21+) does this in place.
- Comparing to `xs[i-1]` (source) vs the last kept value is equivalent here since
  runs are contiguous.
