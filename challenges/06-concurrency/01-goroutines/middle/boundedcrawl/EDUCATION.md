# Bounded Sitemap Crawl

## Intuition

A buffered channel with capacity N is a permit box holding N permits. Take one before you start, put it back when you finish, and at no instant can more than N workers be running — no counters, no condition variables, no lock.

## Approach

1. Allocate `codes := make([]int, len(urls))`.
2. Build `slots := make(chan struct{}, maxParallel)` only when `maxParallel > 0`; leave it nil otherwise.
3. For each URL: `wg.Add(1)`, acquire a slot if the semaphore exists, then launch the goroutine.
4. Inside, `defer wg.Done()` and `defer func() { <-slots }()`, then write `fetch(url)` into `codes[i]`.
5. `wg.Wait()` and return `codes`.

## Solution

```go
// Package boundedcrawl — Gopher Workplace challenge.
package boundedcrawl

import "sync"

// CrawlPages fetches every URL and returns the status codes in input order,
// never running more than maxParallel fetches at the same time. The politeness
// limit protects the origin from being knocked over by its own crawler; a
// maxParallel of zero or less means "no limit".
//
// Examples:
//
//	CrawlPages([]string{"/", "/gone"}, 2, fetch)  => [200 404]
//	CrawlPages([]string{"/"}, 1, fetch)           => [200]
//	CrawlPages(nil, 4, fetch)                     => []
func CrawlPages(urls []string, maxParallel int, fetch func(url string) int) []int {
	codes := make([]int, len(urls))

	var slots chan struct{}
	if maxParallel > 0 {
		slots = make(chan struct{}, maxParallel)
	}

	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		if slots != nil {
			slots <- struct{}{}
		}
		go func(i int, url string) {
			defer wg.Done()
			if slots != nil {
				defer func() { <-slots }()
			}
			codes[i] = fetch(url)
		}(i, url)
	}
	wg.Wait()
	return codes
}
```

## Walkthrough

- With `maxParallel == 1` the send blocks until the previous goroutine's deferred receive runs, so the crawl is effectively sequential and the peak counter never exceeds 1.
- With nine URLs and a limit of three the main goroutine blocks on the fourth send until a slot frees, keeping the peak at three.
- A limit larger than the input never blocks — the buffer is never full.
- With `maxParallel == 0` no channel is created and every URL is fetched at once; the results are still in order.

## Pitfalls

- Acquiring the slot *inside* the goroutine: the limit then caps only concurrent fetches, not the millions of goroutines you already spawned.
- Releasing with a plain `<-slots` at the end of the body instead of a `defer` — one early return and the semaphore leaks a permit forever.
- Sending on a nil channel to represent "unlimited": that blocks forever. Guard the nil case explicitly.
- Sizing the buffer from `len(urls)` instead of `maxParallel`, which quietly removes the limit.
