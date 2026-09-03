# Bounded Sitemap Crawl

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

The sitemap crawler walks tens of thousands of URLs. Spawning one goroutine per URL is cheap in Go but not free at the other end: the origin has a connection budget, and an unbounded fan-out turns your own crawler into a denial-of-service. A counted semaphore caps how many fetches are in flight while still keeping the results in sitemap order.

## Task

Implement the exported function(s) in [boundedcrawl.go](boundedcrawl.go) so that:

1. Return a slice of status codes the same length as `urls`, in input order.
2. Fetch each URL in its own goroutine, joined with a `sync.WaitGroup`.
3. Never run more than `maxParallel` fetches concurrently; a value `<= 0` means unbounded.
4. Release the slot on every exit path, so one slow page cannot starve the rest.
5. A nil or empty URL list returns an empty slice.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  CrawlPages([]string{"/", "/gone"}, 2, fetch)
Output: [200 404]
```

**Example 2:**

```
Input:  CrawlPages([]string{"/a", "/b", "/c"}, 1, fetch)
Output: [200 200 200]
```

**Example 3:**

```
Input:  CrawlPages(nil, 4, fetch)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counting semaphore** | A buffered `chan struct{}` of capacity N: a send takes a slot, a receive gives it back. |
| 2 | **Bounded spawning** | Acquire the slot *before* the `go` statement so the number of live goroutines is capped too. |
| 3 | **`defer` for release** | Releasing in a `defer` guarantees the slot comes back even on an early return or a panic. |
| 4 | **Per-index results** | The semaphore controls parallelism; the index still controls order. |

## Hint

`slots <- struct{}{}` before `go`, and `defer func() { <-slots }()` inside. Skip the channel entirely when `maxParallel <= 0`.

## Validate

```bash
make verify
```
