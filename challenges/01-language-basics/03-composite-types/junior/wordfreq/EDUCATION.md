# Frequency counting with maps

## The idea

The map zero value makes counting a one-liner: a missing key reads as 0, so
`m[k]++` reads-modifies-writes without a presence check:

```go
m := make(map[string]int)
for _, w := range words { m[w]++ }
```

## Why it matters

Histograms, tallies, and grouping all rely on this. It's the canonical use of the
zero-value read.

## Watch out

- `make` the map first; `m[k]++` on a nil map panics (it's a write).
- Iteration order is random — sort keys for deterministic output.
- For non-zero defaults, use comma-ok instead of the zero value.
