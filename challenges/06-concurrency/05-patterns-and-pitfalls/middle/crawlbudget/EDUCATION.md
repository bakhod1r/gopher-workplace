# Bounded Parallelism with a Semaphore

## Intuition

A worker pool bounds concurrency by having a fixed number of consumers. A
semaphore bounds it from the other side: any number of goroutines may exist in
the code, but only `limit` of them may hold a slot at a time. In Go the
semaphore *is* a buffered channel — `sem <- struct{}{}` blocks once the buffer
is full, and `<-sem` frees a place.

The detail that decides whether the budget is real: **acquire before you
spawn**. If you start the goroutine first and acquire inside it, you have
bounded the concurrent *fetches* but created `len(urls)` goroutines, each
holding its closure alive. Acquiring in the dispatch loop makes the loop itself
apply back-pressure, which is what the site owner actually asked for.

The rest is the same shape as any bounded batch: index-keyed results for order,
a mutex-guarded first error, and a derived context so that error stops the
dispatch loop from acquiring even one more slot.

## Approach

1. Guard the dead context, the empty frontier and `limit < 1`.
2. Derive `runCtx, cancel`; `defer cancel()`.
3. Allocate `pages` and `sem := make(chan struct{}, limit)`.
4. In a labeled loop, `select` on acquiring a slot versus `runCtx.Done()`,
   breaking the loop on the latter.
5. Spawn the fetch goroutine with `defer wg.Done()` and `defer func(){ <-sem }()`.
6. Store `pages[i]`, or record the first error and `cancel()`.
7. `wg.Wait()`, then return the error or the ordered pages.

## Solution

```go
// CrawlPages fetches every URL with at most limit fetches in flight at once.
// A buffered channel is used as a counting semaphore: acquiring a slot is a
// send, releasing it is a receive, and the crawl can never exceed the politeness
// budget the site owner agreed to.
//
// Pages come back in the order the URLs were given. The first fetch error
// cancels the crawl and is returned with a nil slice.
//
// Examples:
//
//	CrawlPages(live ctx, 6 urls, 2, fetch)        => 6 pages in order
//	CrawlPages(live ctx, urls with "bad", 3, get) => errFetch
//	CrawlPages(cancelled ctx, urls, 2, fetch)     => context.Canceled
func CrawlPages(ctx context.Context, urls []string, limit int, fetch func(context.Context, string) (Page, error)) ([]Page, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, nil
	}
	if limit < 1 {
		limit = 1
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	pages := make([]Page, len(urls))
	sem := make(chan struct{}, limit)

	var mu sync.Mutex
	var firstErr error
	var wg sync.WaitGroup

dispatch:
	for i, url := range urls {
		select {
		case sem <- struct{}{}:
		case <-runCtx.Done():
			break dispatch
		}

		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			defer func() { <-sem }()

			page, err := fetch(runCtx, url)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
					cancel()
				}
				mu.Unlock()
				return
			}
			pages[i] = page
		}(i, url)
	}

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return pages, nil
}
```

## Walkthrough

- **Four URLs, `limit` 2.** The loop acquires two slots and spawns two fetches;
  the third `sem <- struct{}{}` blocks until one of them runs its `defer <-sem`.
  At no instant do three fetches overlap — which is exactly what the tracker in
  the test asserts, for limits 1, 2 and 5 over 40 URLs.
- **`limit` 1.** The semaphore has one slot, so the crawl is effectively serial,
  yet the code is unchanged. A budget of one is a legitimate configuration, not
  a special case.
- **`bad-1` in the frontier.** Its goroutine takes the mutex, stores `errFetch`
  and cancels. The dispatch loop's next `select` may still win the `sem` arm
  once — that is fine and harmless — but as soon as the buffer is full it takes
  `<-runCtx.Done()` and breaks out. `wg.Wait()` drains the in-flight fetches and
  `errFetch` is returned with no partial page list.
- **Cancelled caller.** Nothing is dispatched at all.

## Pitfalls

- Acquiring the slot *inside* the goroutine bounds the fetches but not the
  goroutines; with a large frontier that is a memory problem.
- `break` inside the `select` breaks the `select`, not the `for`. Without the
  `dispatch:` label the loop keeps going after cancellation.
- Releasing the slot with a plain `<-sem` at the end of the goroutine instead of
  a `defer` leaks a slot on every early return, and the crawl deadlocks once
  `limit` slots have leaked.
- Appending pages to a shared slice under the mutex gives completion order and
  makes the mutex a bottleneck. `pages[i] = page` needs no lock at all.
- Sizing the semaphore with `make(chan struct{}, limit)` but sending `limit + 1`
  times before any release — for example by acquiring twice per URL — is a
  self-inflicted deadlock; each URL acquires exactly once.
