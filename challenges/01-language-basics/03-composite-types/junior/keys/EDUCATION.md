# Deterministic order from a map

## The idea

Go randomizes map iteration order on purpose. To get a stable order, collect the
keys into a slice and sort:

```go
out := make([]string, 0, len(m))
for k := range m { out = append(out, k) }
sort.Strings(out)
```

## Why it matters

Any output that must be reproducible — logs, JSON, tests, hashing — needs an
explicit sort, because relying on map order gives flaky, non-deterministic
results.

## Watch out

- Pre-size with `make(..., 0, len(m))` to avoid regrowth.
- Ranging a map with only `k` iterates keys.
- `slices.Sorted(maps.Keys(m))` (Go 1.23+) does this in one line.
