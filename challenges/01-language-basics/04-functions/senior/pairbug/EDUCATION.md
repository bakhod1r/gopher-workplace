# Look-ahead loop bounds

## Intuition

When the body reads `xs[i+1]`, the loop must terminate at `len(xs)-1`; the classic `< len(xs)` bound over-runs by one.

## Approach

1. Summing adjacent differences accesses `xs[i+1]`, so `i` must stop at `len-2`.
2. The bug loops to `len-1`, indexing out of range; use `i < len(xs)-1`.

## Solution

```go
func SumDiffs(xs []int) int {
	total := 0
	for i := 0; i < len(xs)-1; i++ {
		total += xs[i+1] - xs[i]
	}
	return total
}
```

## Walkthrough

At the last index, `xs[i+1]` is out of bounds and panics. Stopping one early keeps every `xs[i+1]` valid.

## Pitfalls

- Reading `xs[i+1]` ⇒ stop at `len(xs)-1`.
- With 0 or 1 elements the loop must not run.
