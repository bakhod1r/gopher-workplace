# Copying and merging maps

## The idea

Maps are reference types, so returning a fresh, independent map means allocating
one and copying entries in override order:

```go
out := make(map[string]int)
for k, v := range a { out[k] = v }
for k, v := range b { out[k] = v } // b wins on collisions
```

## Why it matters

Layered configuration and defaults+overrides are everyday patterns. Mutating an
input map instead of copying causes spooky action-at-a-distance, because the
caller shares the same underlying map.

## Watch out

- Assigning one map to another copies the **reference**, not the data.
- Iteration order over a map is randomized — don't rely on it.
- Writing to a nil map panics; `make` first.
