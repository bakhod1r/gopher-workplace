# Group-by-sum

## The idea

Accumulate a running total per key; the map zero value makes `+=` work for new
keys:

```go
m := make(map[string]int)
for _, o := range orders { m[o.Customer] += o.Amount }
```

## Why it matters

This is the SQL `GROUP BY ... SUM(...)` in Go. Aggregating records by a key field
is a foundational reporting operation.

## Watch out

- `make` the map first; `+=` is a write.
- The zero-value read is what lets a new key start at 0.
- For multiple aggregates, use a struct as the map value.
