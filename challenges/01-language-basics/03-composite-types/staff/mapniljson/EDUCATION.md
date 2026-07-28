# Nil and empty maps at the boundary

## The idea

Like slices, a nil map and an empty map both have length 0 but serialize
differently: `encoding/json` emits `null` for nil and `{}` for an allocated empty
map:

```go
m := make(map[string]int) // marshals to {}, even when empty
```

## Why it matters

API responses distinguish `{}` (present, empty) from `null` (absent). Returning a
lazily-allocated map means empty input yields `null`, a contract violation that
only shows on the empty path.

## Watch out

- `m == nil` is the only way to tell nil from empty.
- Reading nil maps is safe; writing panics — allocate before the loop.
- Prefer `make` at the top when the emptiness must be observable.
