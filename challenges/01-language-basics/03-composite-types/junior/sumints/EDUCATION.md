# Ranging over slices

## The idea

`for i, v := range xs` visits each element; use `_` for an index you don't need.
Accumulate into a variable declared before the loop:

```go
total := 0
for _, x := range xs { total += x }
```

## Why it matters

Iterate-and-accumulate is the shape of countless slice operations (sum, max,
count, build). Ranging works identically on nil and empty slices — both have
length 0 — so no special case is needed.

## Watch out

- A `nil` slice is safe to range (zero iterations); no need to check.
- The range value is a **copy**; assigning to it doesn't change the slice.
- `len(xs)` is O(1); indexing `xs[i]` is fine too, but range reads cleaner.
