# Streaming Reductions

## Intuition

A reduction folds a stream into one value in a single pass. The subtle part
is initialisation: any constant seed you pick can be wrong for some input.
Seeding from the *first element* is always correct, and the `seen` flag is
what lets you detect that first element.

## Approach

1. Declare `peak` (0) and `seen` (false).
2. `range` over `samples`.
3. If `!seen` or `v > peak`, set `peak = v` and `seen = true`.
4. Return `peak, seen`.

## Solution

```go
func PeakLatency(samples <-chan int) (int, bool) {
	peak := 0
	seen := false
	for v := range samples {
		if !seen || v > peak {
			peak = v
			seen = true
		}
	}
	return peak, seen
}
```

## Walkthrough

For `-5, -2`: the first sample sets `peak = -5` because `!seen`; then
`-2 > -5` updates it to `-2`. Result `-2, true`. Had `peak` started at `0`,
the answer would wrongly have been `0`.

## Pitfalls

- Seeding `peak` with `0` breaks on an all-negative window.
- Collecting into a slice first works but wastes memory — a reduction needs none.
- Returning `peak != 0` instead of `seen` misreports a window of all-zero samples.
