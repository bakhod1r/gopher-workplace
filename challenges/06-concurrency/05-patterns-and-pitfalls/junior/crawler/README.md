# Bounded Crawler

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

A crawler that opens one connection per URL will be blocked by the target site
within seconds. The polite version keeps a fixed number of requests in flight
and reports the status code for each URL *in the order the frontier listed
them*, so the report lines up with the input.

## Task

Implement `CrawlSite` in [crawler.go](crawler.go) so that:

1. It preallocates the result slice with `make([]int, len(urls))`.
2. It starts one goroutine per URL, each gated by a semaphore of capacity `limit`.
3. Each goroutine writes `fetch(u)` into its own index, and the function returns the slice after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CrawlSite([]string{"/a", "/bb"}, 2, fetch)
Output: []int{2, 3}
```

**Example 2:**

```
Input:  CrawlSite([]string{"/aaaa", "/b", "/ccc"}, 2, fetch)
Output: []int{5, 2, 4}
```

**Example 3:**

```
Input:  CrawlSite(nil, 4, fetch)
Output: []int{} (empty, length 0)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded parallelism** | Semaphore capacity, not the frontier size, decides peak concurrency. |
| 2 | **Index-keyed writes** | Distinct slice elements can be written concurrently without a mutex. |
| 3 | **Preallocation** | `append` from goroutines would race; a fixed-length slice does not. |

## Hint

Allocate the result slice *before* starting any goroutine, and pass both `i`
and the URL into the goroutine so each one writes its own slot.

## Validate

```bash
make verify
```
