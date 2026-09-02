# Bounded Parallelism with Ordered Results

## Intuition

Two goroutines writing `codes[0]` and `codes[1]` touch different memory, so
there is no race — the race detector agrees. That is why index-keyed writes
are the standard way to keep concurrent results in input order.

## Approach

1. `codes := make([]int, len(urls))`.
2. For each `i, u`: goroutine acquires a permit, writes `codes[i] = fetch(u)`, releases.
3. `wg.Wait()` then return `codes`.

## Solution

```go
import "sync"

func CrawlSite(urls []string, limit int, fetch func(string) int) []int {
	codes := make([]int, len(urls))
	sem := make(chan struct{}, limit)

	var wg sync.WaitGroup
	for i, u := range urls {
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			codes[i] = fetch(u)
		}(i, u)
	}

	wg.Wait()
	return codes
}
```

## Walkthrough

For three URLs with limit 2, the third goroutine waits for a permit, but
whenever it runs it writes index 2. Execution order varies run to run; the
returned slice always matches the frontier order.

## Pitfalls

- Using `codes = append(codes, fetch(u))` from goroutines: a race on the slice header, and scrambled order.
- Spawning one goroutine per URL with no semaphore — that is exactly the unbounded spawn that gets you blocked.
- Reading the slice before `wg.Wait()` returns — the writes are not guaranteed visible yet.
