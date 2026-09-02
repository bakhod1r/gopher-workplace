# Price Fetcher

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A storefront renders a category page. Every SKU on the page needs a live price
from the pricing service, and the calls are independent, so the page fetches
them all at once instead of one after another. The store discount is then
applied to each price.

## Task

Implement `FetchAllPrices` in [pricefetcher.go](pricefetcher.go) so that:

1. Preallocate the result with `make([]int, len(skus))`.
2. Start one goroutine per SKU; goroutine `i` writes `fetch(skus[i]) * (100 - discountPct) / 100` to `out[i]`.
3. Wait for every goroutine with a `sync.WaitGroup` before returning.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FetchAllPrices([]string{"ab", "cde"}, catalog, 0)
Output: [200 300]
```

**Example 2:**

```
Input:  FetchAllPrices([]string{"ab"}, catalog, 50)
Output: [100]
```

**Example 3:**

```
Input:  FetchAllPrices(nil, catalog, 10)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Fan-out over I/O** | Independent remote calls are the classic reason to reach for goroutines. |

## Hint

`fetch` is injected so the tests stay deterministic. Pass `i` and `sku` into the
goroutine; `fetch` and `discountPct` are fixed for the whole call, so capturing
them is safe.

## Validate

```bash
make verify
```
