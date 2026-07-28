# Zipping parallel slices

## The idea

Iterate to the shorter length and pair by index:

```go
n := len(keys)
if len(vals) < n { n = len(vals) }
for i := 0; i < n; i++ { m[keys[i]] = vals[i] }
```

## Why it matters

Parallel arrays (headers/row, ids/names) are common at boundaries; zipping turns
them into a keyed structure. Guarding on the shorter length avoids index panics.

## Watch out

- Mismatched lengths: iterate `min`, don't assume equal.
- Duplicate keys: later pairs overwrite earlier ones.
- `min` is a builtin in Go 1.21+.
