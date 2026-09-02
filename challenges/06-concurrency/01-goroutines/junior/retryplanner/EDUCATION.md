# Retry Planner

## Intuition

A loop with an accumulator is the ideal goroutine body: everything it touches is
local, and only the final value escapes through a single indexed write.

## Approach

1. Allocate `out := make([]int, len(attempts))`.
2. In each goroutine, set `delay := baseMs` and double it `n` times.
3. Store `out[i] = delay`, then `wg.Wait()`.

## Solution

```go
// Package retryplanner — Gopher Workplace challenge.
package retryplanner

import (
	"sync"
)

// Backoffs returns the exponential backoff delay for each attempt number.
//
// Examples:
//
//	Backoffs([]int{0, 1, 3}, 100)  => [100 200 800]
//	Backoffs([]int{-1}, 100)       => [100]
//	Backoffs(nil, 100)             => []
func Backoffs(attempts []int, baseMs int) []int {
	out := make([]int, len(attempts))
	var wg sync.WaitGroup
	for i, n := range attempts {
		wg.Add(1)
		go func(i int, n int) {
			defer wg.Done()
			delay := baseMs
			for k := 0; k < n; k++ {
				delay *= 2
			}
			out[i] = delay
		}(i, n)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- Attempt `0` has not failed yet, so the delay is the base `100`.
- Attempt `3` doubles three times: `100 → 200 → 400 → 800`.
- A negative attempt number never enters the loop and also yields the base.

## Pitfalls

- Declaring `delay` outside the goroutine, which makes every schedule racy and wrong.
- Starting the accumulator at `0`, so every delay stays `0`.
- Special-casing negatives when the loop already handles them.
