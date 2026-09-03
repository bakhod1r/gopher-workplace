# Thumbnail Worker Pool

**Level:** middle
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The media service renders thumbnails for every upload in a batch. One worker
is too slow and one goroutine per upload melts the box, so the batch is drained
by a fixed pool. Two things make this harder than a plain fan-out: the caller
needs the thumbnails **back in upload order**, and the first render failure
should stop the pool rather than let it chew through the rest of the batch.

## Task

Implement `RenderQueue` in [thumbnailworkers.go](thumbnailworkers.go) so that:

1. It returns `ctx.Err()` immediately if the caller's context is already finished, and `nil, nil` for an empty queue.
2. It starts exactly `workers` goroutines pulling indexes off one jobs channel (`workers < 1` is treated as 1).
3. Each result is written to `out[i]`, so the returned slice matches the input order without sorting.
4. The first render error is remembered, cancels a derived context so the feeder and the other workers stop, and is returned with a `nil` slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RenderQueue(live ctx, [a bb ccc], 2 workers, render)
Output: [a.thumb bb.thumb ccc.thumb], nil
```

**Example 2:**

```
Input:  RenderQueue(live ctx, [a bad-1 c], 3 workers, render)
Output: nil, errRender
```

**Example 3:**

```
Input:  RenderQueue(cancelled ctx, [a b], 2 workers, render)
Output: nil, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Worker pool** | N goroutines ranging over one jobs channel; closing it retires the pool. |
| 2 | **Index-keyed results** | Writing to distinct `out[i]` slots is race-free and preserves order — no mutex, no sort. |
| 3 | **Error-triggered cancellation** | A derived context turns one worker's failure into a stop signal for the whole pool. |
| 4 | **Feeder select** | The feeder sends on `jobs` *or* observes `ctx.Done()`, so it can never block after a cancel. |

## Hint

Send indexes, not uploads, down the jobs channel — then `out[i] = thumb`
needs no lock. Guard the shared `firstErr` with a mutex and call `cancel()`
the first time it is set.

## Validate

```bash
make verify
```
