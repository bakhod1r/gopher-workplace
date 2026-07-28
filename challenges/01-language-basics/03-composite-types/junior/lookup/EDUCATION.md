# Map lookup

## The idea

Indexing a map returns the value, or the value type's zero if the key is absent.
The comma-ok form reports presence:

```go
v, ok := m[key]
```

## Why it matters

Maps are the go-to associative structure. Distinguishing "absent" from "present
with zero value" requires the `ok` result, not the value.

## Watch out

- Reading a nil map is safe (returns zero); writing panics.
- Iteration order is randomized.
- Keys must be comparable.
