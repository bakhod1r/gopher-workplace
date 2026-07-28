# Slice expressions

## The idea

`xs[low:high]` creates a slice over the same backing array covering `[low, high)`.
It is O(1) and shares memory with the original:

```go
mid := xs[1:4] // elements 1,2,3
```

## Why it matters

Sub-slicing is how you window, page, and split data cheaply. But because it
shares memory, writes through a sub-slice affect the original.

## Watch out

- High index must be ≤ `len` (or `cap` for the two-index form) or it panics.
- Sub-slices share the backing array; append may or may not alias depending on
  capacity.
- The three-index form `xs[a:b:c]` caps capacity.
