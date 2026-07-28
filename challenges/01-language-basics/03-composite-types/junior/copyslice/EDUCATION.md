# Slices share backing arrays

## The idea

A slice is a header (pointer, length, capacity) over a backing array. Copying the
header (`b := a`) shares the array. An independent copy allocates a new array:

```go
out := make([]int, len(xs))
copy(out, xs)
```

## Why it matters

Aliasing bugs — mutating a "copy" that secretly shares memory — are among the
most common slice mistakes. `make`+`copy` (or `slices.Clone`) gives independence.

## Watch out

- `copy` copies `min(len(dst), len(src))` elements.
- `make([]int, len(xs))` is non-nil even when `len` is 0.
- `append` to a shared slice may alias depending on capacity.
