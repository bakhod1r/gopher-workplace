# Structs as map keys

## The idea

A struct whose fields are all comparable is itself comparable and can be a map
key. Its identity is the tuple of field values, so order/assignment matters:

```go
m[Point{p.X, p.Y}]++ // or simply m[p]++
```

## Why it matters

Composite keys (coordinates, versioned IDs, small records) are cleanly expressed
as structs — no string concatenation, no ambiguity. Building the key with swapped
or wrong fields silently merges distinct keys.

## Watch out

- A struct with a slice/map/func field is **not** comparable — compile error as a
  key.
- `m[p]` reuses the loop value directly, avoiding field mistakes.
- Keys are hashed/compared by all fields.
