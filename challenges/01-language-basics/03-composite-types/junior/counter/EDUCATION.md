# Counting with the map zero value

## The idea

Reading a missing map key yields the value type's zero — 0 for ints. So
incrementing works without checking presence:

```go
m := make(map[string]int)
for _, x := range xs { m[x]++ }
```

`m[x]++` reads (0 if absent), adds one, and stores.

## Why it matters

Histograms, tallies, and grouping all lean on this. The zero-value read is what
makes `m[k]++` a clean one-liner instead of a presence check plus insert.

## Watch out

- You must `make` the map first; `m[x]++` on a nil map panics (it's a write).
- Only counts are convenient; for non-zero defaults you still need comma-ok.
- Iteration order of the result is random.
