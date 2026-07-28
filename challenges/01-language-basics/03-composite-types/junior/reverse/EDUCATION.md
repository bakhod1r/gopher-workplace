# Reversing a slice

## The idea

Swap symmetric elements from the ends inward, stopping at the middle:

```go
for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 { xs[i], xs[j] = xs[j], xs[i] }
```

## Why it matters

Reversal is a building block (e.g. rotate = three reversals) and demonstrates the
two-pointer technique and in-place mutation of a slice's backing array.

## Watch out

- In-place reversal mutates the caller's slice (shared backing array).
- Multiple assignment swaps without a temp.
- `slices.Reverse` does this generically (Go 1.21+).
