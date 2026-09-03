# ns/op That Drifts With `b.N`

## Intuition

A benchmark reports elapsed time divided by the iteration count. Any fixed cost inside that elapsed time gets divided too, so it contributes `setup/N` to ns/op — a term that shrinks as the harness raises `N`.

## Approach

1. Drop the setup term entirely.
2. Return the per-iteration work multiplied by the iteration count.

## Solution

```go
func Measured(setupNS, workNS, n int64) int64 {
	if n <= 0 {
		return 0
	}
	return workNS * n
}
```

## Walkthrough

With the bug, `PerOp(1_000_000, 7, 10)` is `100_007` and `PerOp(1_000_000, 7, 1_000_000)` is `8` — the same code, the same machine, a four-orders-of-magnitude difference driven purely by how many iterations the harness chose to run.

## Pitfalls

- Amortising the setup as `setupNS/n` instead of removing it; the result is still `N`-dependent.
- Calling `b.ResetTimer()` after the loop, which discards the measurement instead of the setup.
- Building the fixture inside the loop, where it becomes part of every iteration.
