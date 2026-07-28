# Slicing array pointers

## The idea

`(&arr)[:]` produces a slice over the entire array, aliasing it; a shorter slice expression drops elements.

## Why it matters

Exposing a fixed array as a slice view is a common zero-copy adapter.

## Watch out

- `p[:]` gives all 4 elements; `p[:2]` only the first two.
- The resulting slice aliases the array's memory.
