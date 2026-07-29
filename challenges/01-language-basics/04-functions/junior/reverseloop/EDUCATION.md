# Index-based reversal

## Intuition

Reading source back-to-front into a new slice keeps the original intact — a copy, not an in-place swap.

## Approach

1. Loop from the last index down to 0.
2. Append each element to a fresh slice.

## Solution

```go
func Reverse(xs []int) []int {
	out := make([]int, 0, len(xs))
	for i := len(xs) - 1; i >= 0; i-- {
		out = append(out, xs[i])
	}
	return out
}
```

## Walkthrough

`Reverse([1 2 3])` appends 3, 2, 1 in that order.

## Pitfalls

- Reversing in place would mutate the caller's slice (shared backing array).
- Start the index at `len(xs)-1`, end at `>= 0`.
