# Indexing needs length, appending needs capacity

## The idea

`make([]int, 0, n)` reserves capacity but has length 0 — indexing `out[i]` panics.
Two valid patterns:

```go
out := make([]int, len(xs)); for i, x := range xs { out[i] = x * 2 } // index
out := make([]int, 0, len(xs)); for _, x := range xs { out = append(out, x*2) } // append
```

## Why it matters

Confusing the two forms is a frequent slice bug: you pre-size "for efficiency"
with capacity, then index into nonexistent elements. Pick index **or** append and
size accordingly.

## Watch out

- Capacity is spare room; only elements within length exist.
- `append` grows length; indexing requires it to already be there.
- `make([]T, n)` zero-fills n elements; `make([]T, 0, n)` fills none.
