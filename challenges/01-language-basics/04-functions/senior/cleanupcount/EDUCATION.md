# Defer runs at function scope, not block scope

## Intuition

A `defer` inside a loop is deferred to the enclosing function's return, so resources accumulate until then — sometimes the intent, sometimes a leak.

## Approach

1. The peak must be observed before the counter drains.
2. The bug decrements inline during the loop; defer the decrement so draining happens after the peak is recorded.

## Solution

```go
func PeakThenDrain(n int) (peak int) {
	open := 0
	for i := 0; i < n; i++ {
		open++
		if open > peak {
			peak = open
		}
		defer func() { open-- }()
	}
	return peak
}
```

## Walkthrough

Decrementing `open` inline lowers the peak before it's captured. Deferring `open--` postpones draining until after the function records the maximum.

## Pitfalls

- Defers do not run at each iteration's end.
- For per-iteration cleanup, extract the body into its own function.
