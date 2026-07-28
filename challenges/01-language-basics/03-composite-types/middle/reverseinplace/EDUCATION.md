# In-place reversal

## The idea

Swap symmetric pairs from the outside in, stopping at the middle:

```go
for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
	xs[i], xs[j] = xs[j], xs[i]
}
```

## Why it matters

In-place algorithms avoid allocation — important in hot paths. Reversal is also a
building block: rotating a slice is three reversals.

## Watch out

- This **mutates** the caller's slice (shared backing array); copy first if the
  original must survive.
- Multiple assignment swaps without a temp.
- `slices.Reverse` (Go 1.21+) does this generically.
