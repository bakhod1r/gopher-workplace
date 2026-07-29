# Full-slice expressions and retention

## Intuition

`s[:k:k]` limits capacity so appends reallocate instead of spilling; it also lets the unused tail be collected once the parent is gone.

## Approach

1. `xs[:k]` keeps the original capacity, so appending overwrites `xs[k]`.
2. Use a full-slice expression `xs[:k:k]` to cap capacity at k.

## Solution

```go
func Prefix(xs []int, k int) []int {
	return xs[:k:k]
}
```

## Walkthrough

With `xs[:2]` the capacity is still 5, so `append` writes into `xs[2]`. `xs[:2:2]` forces a reallocation on append, protecting `xs`.

## Pitfalls

- `xs[:k]` keeps cap == cap(xs); appends spill into the parent.
- `xs[:k:k]` (or a copy) isolates the prefix.
