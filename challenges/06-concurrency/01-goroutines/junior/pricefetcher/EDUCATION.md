# Price Fetcher

## Intuition

Fetching one price tells you nothing about the next, so the page can issue all
the calls at once. The result still has to line up with the input order, and it
does — because position comes from the index, not from which call returns first.

## Approach

1. Allocate the output slice so every index exists before any goroutine runs.
2. `wg.Add(1)` before each `go`, `defer wg.Done()` inside the goroutine.
3. Each goroutine writes exactly one slot: `out[i] = fetch(sku) * (100 - discountPct) / 100`.
4. `wg.Wait()` blocks until all writes are done, then return `out`.

## Solution

```go
// Package pricefetcher — Gopher Workplace challenge.
package pricefetcher

import (
	"sync"
)

// FetchAllPrices returns the discounted price of every SKU, in input order.
//
// Examples:
//
//	FetchAllPrices([]string{"ab", "cde"}, catalog, 0)   => [200 300]
//	FetchAllPrices([]string{"ab"}, catalog, 50)         => [100]
//	FetchAllPrices(nil, catalog, 10)                    => []
func FetchAllPrices(skus []string, fetch func(sku string) int, discountPct int) []int {
	out := make([]int, len(skus))
	var wg sync.WaitGroup
	for i, sku := range skus {
		wg.Add(1)
		go func(i int, sku string) {
			defer wg.Done()
			out[i] = fetch(sku) * (100 - discountPct) / 100
		}(i, sku)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `{"ab", "cde"}` with a 0% discount starts two goroutines.
- They write `200` and `300` to indices 0 and 1, in whatever order they finish.
- With a 50% discount, `200 * 50 / 100` is `100`.

## Pitfalls

- Collecting prices with `append` from goroutines — that races and scrambles the page order.
- Calling `wg.Add(1)` inside the goroutine, so `Wait` may return before it starts.
- Returning `out` before `wg.Wait()`, handing the caller a slice full of zeros.
