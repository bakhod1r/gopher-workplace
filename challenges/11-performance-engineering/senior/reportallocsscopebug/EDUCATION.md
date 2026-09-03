# Dividing By The Wrong Iterations

## Intuition

A per-operation number is a ratio, and a ratio is only meaningful when both halves describe the same thing. Here the counters cover the measured phase and the divisor covers the whole run.

## Approach

1. Divide by the measured iteration count alone.

## Solution

```go
func PerOp(r Run) (uint64, uint64) {
	if r.Measured <= 0 {
		return 0, 0
	}
	n := uint64(r.Measured)
	return r.Bytes / n, r.Allocs / n
}
```

## Walkthrough

With the bug and a 90-iteration warmup, 20 allocations over 10 measured iterations are divided by 100 and reported as 0 allocs/op — and the more warmup you add, the more allocation-free the function appears.

## Pitfalls

- Resetting the timer but not the memory counters, which is the mirror-image bug.
- Reporting a `0 allocs/op` result without checking the divisor.
- Comparing per-op numbers between runs with different warmup strategies.
