# Reverse lookups

## The idea

Inverting a map swaps roles: each `key→value` becomes `value→key`. The result
map's types are the originals reversed:

```go
out := make(map[int]string)
for k, v := range m { out[v] = k }
```

## Why it matters

Bidirectional lookups (id↔name, code↔label) are common. Building the inverse once
is cheaper than scanning for a value every time.

## Watch out

- If values are **not** unique, later pairs overwrite earlier ones — inversion
  loses data.
- The value type must be a valid map key (comparable).
- Iteration order is random; the resulting map is unordered.
