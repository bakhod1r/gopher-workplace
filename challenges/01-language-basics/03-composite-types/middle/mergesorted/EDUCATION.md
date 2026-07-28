# The merge step

## The idea

Two indices walk the sorted inputs; append the smaller front element, advance
that side, and finally drain whatever remains:

```go
i, j := 0, 0
for i < len(a) && j < len(b) {
	if a[i] <= b[j] { out = append(out, a[i]); i++ } else { out = append(out, b[j]); j++ }
}
out = append(out, a[i:]...)
out = append(out, b[j:]...)
```

## Why it matters

Merging sorted runs is the heart of merge sort and external sorting. It's O(n+m),
far cheaper than concatenate-then-sort.

## Watch out

- Use `<=` (not `<`) to keep the merge stable for equal keys.
- Remember to drain both tails — one will be empty.
- Pre-size the output to `len(a)+len(b)` to avoid regrowth.
