# Sitemap Crawler

## Intuition

Crawling is the textbook fan-out. The one thing you must not do is have
goroutines append to a shared report — index them instead, and both the race and
the ordering problem vanish together.

## Approach

1. Allocate `out := make([]int, len(urls))`.
2. For each `i, url`, `wg.Add(1)` and launch a goroutine that writes `out[i] = get(url)`.
3. `wg.Wait()`, then return `out`.

## Solution

```go
// Package sitemapcrawler — Gopher Workplace challenge.
package sitemapcrawler

import (
	"sync"
)

// FetchStatuses fetches every URL and reports the HTTP status codes in order.
//
// Examples:
//
//	FetchStatuses([]string{"/", "/gone"}, get)  => [200 404]
//	FetchStatuses([]string{"/"}, get)           => [200]
//	FetchStatuses(nil, get)                     => []
func FetchStatuses(urls []string, get func(url string) int) []int {
	out := make([]int, len(urls))
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			out[i] = get(url)
		}(i, url)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"/"` returns `200` and `"/gone"` returns `404`.
- Whichever request completes first, each writes only its own index.
- The report reads `[200 404]` — sitemap order, not response order.

## Pitfalls

- `append`-ing statuses from goroutines: a race, plus a randomly ordered report.
- Reading the loop variable inside the goroutine instead of taking it as a parameter.
- Launching the goroutines but never calling `wg.Wait()`.
