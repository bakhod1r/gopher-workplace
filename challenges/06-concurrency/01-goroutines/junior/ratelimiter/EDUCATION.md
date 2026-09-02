# Rate Limiter

## Intuition

Concurrency is only dangerous where writes are shared. The plan limit is read-
only for the whole call, so no goroutine can observe it changing.

## Approach

1. Allocate `out := make([]int, len(used))`.
2. Launch one goroutine per tenant, passing `i` and the spent count.
3. Subtract, clamp at `0`, and store the result at `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package ratelimiter — Gopher Workplace challenge.
package ratelimiter

import (
	"sync"
)

// RemainingQuota returns each tenant's remaining request quota, never below zero.
//
// Examples:
//
//	RemainingQuota([]int{10, 90}, 100)  => [90 10]
//	RemainingQuota([]int{150}, 100)     => [0]
//	RemainingQuota(nil, 100)            => []
func RemainingQuota(used []int, limit int) []int {
	out := make([]int, len(used))
	var wg sync.WaitGroup
	for i, u := range used {
		wg.Add(1)
		go func(i int, u int) {
			defer wg.Done()
			left := limit - u
			if left < 0 {
				left = 0
			}
			out[i] = left
		}(i, u)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- A tenant that spent 10 of 100 has `90` left.
- A tenant that spent 150 of 100 would be `-50`, so it is clamped to `0`.
- Spending exactly the limit also yields `0`.

## Pitfalls

- Reporting a negative allowance, which the dashboard cannot render.
- Decrementing a single shared counter instead of writing per tenant — a data race.
- Clamping on the parent after `Wait`, which hides the per-tenant rule.
