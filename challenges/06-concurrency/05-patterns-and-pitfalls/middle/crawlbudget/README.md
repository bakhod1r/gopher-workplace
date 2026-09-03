# Crawl Within the Politeness Budget

**Level:** middle
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The crawler has a frontier of URLs and an agreement with the site owner: never
more than `limit` requests in flight. One goroutine per URL would blow that
budget instantly; a fixed worker pool respects it but this crawler also wants
the results back keyed to the frontier order, and a hard stop the moment a
fetch fails. A counting semaphore built from a buffered channel gives you the
budget without a pool.

## Task

Implement `CrawlPages` in [crawlbudget.go](crawlbudget.go) so that:

1. It returns `ctx.Err()` for an already-finished context and `nil, nil` for an
   empty frontier; `limit < 1` is treated as 1.
2. `sem := make(chan struct{}, limit)` is the budget: a goroutine is started
   only after a slot has been acquired, and the slot is released with
   `defer func() { <-sem }()`.
3. Each fetch writes its page to `pages[i]`, so the result matches frontier
   order.
4. The first fetch error is recorded under a mutex, cancels the derived
   context, and is returned with a nil slice. The dispatch loop stops acquiring
   once that context is done.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CrawlPages(live ctx, [a bb ccc dddd], 2, fetch)
Output: pages for a, bb, ccc, dddd in that order
```

**Example 2:**

```
Input:  CrawlPages(live ctx, [a bad-1 c], 2, fetch)
Output: nil, errFetch
```

**Example 3:**

```
Input:  CrawlPages(cancelled ctx, [a b], 2, fetch)
Output: nil, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counting semaphore** | A buffered channel of capacity N: send to acquire, receive to release. |
| 2 | **Acquire before spawn** | Blocking in the dispatch loop is what bounds the goroutine count, not just the work. |
| 3 | **Labeled break in select** | `break` inside a `select` breaks the select; you need a label to leave the loop. |
| 4 | **Release with defer** | A slot leaked on an error path shrinks the budget until the crawl stalls. |

## Hint

Acquire the slot in the loop *before* `go func`, and release it inside the
goroutine with a `defer`. Label the dispatch loop so the `<-runCtx.Done()` arm
can break out of it.

## Validate

```bash
make verify
```
