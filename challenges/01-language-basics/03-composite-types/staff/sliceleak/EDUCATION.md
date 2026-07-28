# Sub-slices pin the whole backing array

## The idea

A slice keeps its backing array reachable for the GC — the *entire* array, not
just the visible window. `xs[:3]` of a 1M-capacity slice keeps all 1M alive:

```go
return append([]int{}, xs[:3]...) // independent; source array can be freed
```

## Why it matters

This is a classic Go memory leak: keeping a tiny slice of a huge buffer (a header,
a token) silently retains megabytes. Copying the small part you need lets the
large array be collected.

## Watch out

- `cap`, not `len`, reveals how much is retained.
- Same leak with sub-strings of huge strings (copy with `strings.Clone`).
- `slices.Clip` reduces capacity but still shares the array; a copy fully
  releases it.
