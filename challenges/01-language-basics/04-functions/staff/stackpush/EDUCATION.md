# Recording per-iteration data in defers

## Intuition

A deferred closure that reads a fixed expression captures none of the loop's progression; snapshot `xs[i]` per iteration to record distinct values.

## Approach

1. Deferred calls run LIFO, reversing the push order.
2. The bug reads `xs[len-1]` at defer-run time (all the same); snapshot the value with a deferred argument `(xs[i])`.

## Solution

```go
func ReverseInts(xs []int) (out []int) {
	for i := 0; i < len(xs); i++ {
		defer func(v int) { out = append(out, v) }(xs[i])
	}
	return
}
```

## Walkthrough

Closing over `xs[len(xs)-1]` makes every deferred call append the last element. Passing `xs[i]` as a deferred argument snapshots each element, so LIFO yields the reversal.

## Pitfalls

- `xs[len(xs)-1]` is the same element every iteration.
- Snapshot `xs[i]` as an argument for a version-proof reverse.
