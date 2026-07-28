# copy respects length, not capacity

## The idea

`copy(dst, src)` copies `min(len(dst), len(src))` elements. A slice made with
`make([]int, 0, n)` has **length 0**, so `copy` writes nothing — capacity is
irrelevant:

```go
dst := make([]int, len(xs)) // length = len(xs)
copy(dst, xs)
```

## Why it matters

The `make(..., 0, n)` + `copy` mistake silently produces empty results. It's a
common confusion between length (how many elements exist) and capacity (room
before regrowth).

## Watch out

- To pre-allocate for `append`, use length 0 + capacity; to `copy`, use full
  length.
- `copy` returns the number copied — useful to assert.
- `slices.Clone` avoids the footgun entirely.
