# The Setup That Should Not Count

## Intuition

A benchmark answers "how long does one operation take". A one-time fixture build is not one operation, so it must not appear in the total.

## Approach

1. Return `0` for a non-positive `n`.
2. Return `workNS * n`; ignore `setupNS` entirely.

## Solution

```go
func Measured(setupNS, workNS int64, n int64) int64 {
	if n <= 0 {
		return 0
	}
	return workNS * n
}
```

## Walkthrough

With setup included, `Measured(1000, 7, 3)` would be `1021` and ns/op would read `340` instead of `7` — a 48x lie that shrinks as `b.N` grows, making the benchmark unstable.

## Pitfalls

- Adding `setupNS` "because it happened".
- Amortizing setup as `setupNS/n`, which still leaves the number `N`-dependent.
- Forgetting the `n <= 0` guard and returning a negative total.
