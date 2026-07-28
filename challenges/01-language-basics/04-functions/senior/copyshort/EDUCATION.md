# copy is bounded by length

## The idea

`copy` moves `min(len(dst), len(src))` elements; a zero-length destination copies nothing regardless of capacity.

## Why it matters

A `make([]T, 0, n)` + `copy` combination is a classic silent no-op clone.

## Watch out

- `copy` uses LENGTH, not capacity, of both slices.
- For a clone use `make([]T, len(src))` (or `append([]T(nil), src...)`).
