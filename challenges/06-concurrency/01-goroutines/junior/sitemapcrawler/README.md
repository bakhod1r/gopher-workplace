# Sitemap Crawler

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A link checker walks a sitemap and requests every URL it lists. The requests do
not depend on one another, so the crawler issues them concurrently and writes
each status code back into the row it came from, keeping the report aligned with
the sitemap.

## Task

Implement `FetchStatuses` in [sitemapcrawler.go](sitemapcrawler.go) so that:

1. Return a slice of status codes the same length as `urls`.
2. Element `i` is `get(urls[i])`.
3. Issue each request in its own goroutine and join them with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FetchStatuses([]string{"/", "/gone"}, get)
Output: [200 404]
```

**Example 2:**

```
Input:  FetchStatuses([]string{"/"}, get)
Output: [200]
```

**Example 3:**

```
Input:  FetchStatuses(nil, get)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Order is not completion order** | A slow URL still lands in its own row, because the index is decided before the goroutine starts. |

## Hint

This is the bare pattern: preallocate, `go` per index, `defer wg.Done()`,
`wg.Wait()`. No shared counters anywhere.

## Validate

```bash
make verify
```
