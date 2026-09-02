# Cert Checker

## Intuition

Each certificate is judged on its own, so the monitor is a pure fan-out. The
boolean result type keeps every goroutine's output to a single slot write.

## Approach

1. Allocate `out := make([]bool, len(expiries))`.
2. Launch one goroutine per certificate, passing `i` and the expiry day.
3. Compute `daysLeft` and store `daysLeft <= window` at `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package certchecker — Gopher Workplace challenge.
package certchecker

import (
	"sync"
)

// ExpiringSoon flags every certificate that expires within the alert window.
//
// Examples:
//
//	ExpiringSoon([]int{100, 400}, 90, 30)  => [true false]
//	ExpiringSoon([]int{50}, 90, 30)        => [true]
//	ExpiringSoon(nil, 90, 30)              => []
func ExpiringSoon(expiries []int, today int, window int) []bool {
	out := make([]bool, len(expiries))
	var wg sync.WaitGroup
	for i, day := range expiries {
		wg.Add(1)
		go func(i int, day int) {
			defer wg.Done()
			daysLeft := day - today
			out[i] = daysLeft <= window
		}(i, day)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- Day `100` against today `90` leaves 10 days, which is inside a 30-day window.
- Day `50` is already 40 days past, so `daysLeft` is negative and the flag is `true`.
- Day `121` leaves 31 days — one day outside the window, so `false`.

## Pitfalls

- Using `<` instead of `<=`, which misses a certificate expiring exactly on the boundary.
- Treating a negative `daysLeft` as safe, so expired certificates are never alerted.
- Appending flagged hosts to a shared slice from goroutines rather than indexing.
