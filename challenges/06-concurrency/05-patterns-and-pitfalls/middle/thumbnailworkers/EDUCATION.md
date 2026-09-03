# Worker Pool with Ordered Results and Error Cancellation

## Intuition

A worker pool is two channels and a WaitGroup: one goroutine feeds jobs, N
workers consume them, and `wg.Wait()` tells you they are all done. Everything
awkward about this puzzle comes from the two extras layered on top.

**Order.** Results arrive in completion order, which is unpredictable. Instead
of collecting them on a results channel and sorting, send *indexes* as the job
and let each worker write to `out[i]`. Distinct slice elements are distinct
memory, so concurrent writes to different indexes are not a data race — and the
slice is already in input order when the pool retires.

**Cancellation.** The first failure should stop the batch. Deriving a
cancellable context inside the function gives you a switch that a worker can
flip and that the feeder can watch. The feeder's `select` is the important
half: without it, the feeder blocks forever on `jobs <- i` once the workers stop
reading, and `wg.Wait()` never returns.

## Approach

1. Reject a finished context up front; return `nil, nil` for an empty queue.
2. Derive `runCtx, cancel` and `defer cancel()` so nothing leaks on any path.
3. Allocate `out` with `len(uploads)` and an unbuffered `jobs chan int`.
4. Start `workers` goroutines: skip the job if `runCtx` is done, otherwise
   render and either store `out[i]` or record the first error and `cancel()`.
5. Start a feeder that sends every index, selecting against `runCtx.Done()`,
   and closes `jobs` on the way out.
6. `wg.Wait()`, then return the first error, the parent's error, or `out`.

## Solution

```go
// RenderQueue drains the upload queue with a pool of worker goroutines and
// returns the thumbnails in the same order as uploads. The pool shares one
// cancellable context: the first render error cancels it, so the remaining
// workers stop pulling new jobs instead of burning CPU on a batch that is
// already doomed.
//
// It returns ctx.Err() if the caller's context is already finished, the first
// render error if any upload failed, or the ordered thumbnails.
//
// Examples:
//
//	RenderQueue(live ctx, 3 uploads, 2 workers, render)   => 3 thumbnails in order
//	RenderQueue(live ctx, uploads with "bad", 2, render)  => errRender
//	RenderQueue(cancelled ctx, 3 uploads, 2, render)      => context.Canceled
func RenderQueue(ctx context.Context, uploads []Upload, workers int, render func(context.Context, Upload) (Thumbnail, error)) ([]Thumbnail, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if len(uploads) == 0 {
		return nil, nil
	}
	if workers < 1 {
		workers = 1
	}

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	out := make([]Thumbnail, len(uploads))
	jobs := make(chan int)

	var mu sync.Mutex
	var firstErr error

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				if runCtx.Err() != nil {
					continue
				}
				thumb, err := render(runCtx, uploads[i])
				if err != nil {
					mu.Lock()
					if firstErr == nil {
						firstErr = err
						cancel()
					}
					mu.Unlock()
					continue
				}
				out[i] = thumb
			}
		}()
	}

	go func() {
		defer close(jobs)
		for i := range uploads {
			select {
			case jobs <- i:
			case <-runCtx.Done():
				return
			}
		}
	}()

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
```

## Walkthrough

- **Happy path, 3 uploads and 2 workers.** The feeder pushes 0, 1, 2; whichever
  worker is free takes the next index. Worker A may finish index 2 before worker
  B finishes index 1, but they wrote to `out[2]` and `out[1]`, so the returned
  slice reads `a.thumb, bb.thumb, ccc.thumb` regardless.
- **`bad-1` in the middle.** Some worker gets `errRender`, takes the mutex, sees
  `firstErr == nil`, stores the error and cancels. The feeder's next send loses
  to `<-runCtx.Done()` and closes `jobs`; the remaining workers drain what is
  already queued but return early on the `runCtx.Err()` guard. `wg.Wait()`
  returns and the function reports `errRender` with a nil slice.
- **Cancelled parent.** The guard at the top fires before a single goroutine is
  started, so the pool never exists.
- **Empty queue.** No feeder work, `jobs` closes immediately, every worker's
  `range` ends at once.

## Pitfalls

- Collecting results on a channel and appending them gives you completion order,
  not input order. Send indexes instead.
- Forgetting the `select` in the feeder deadlocks the whole function after a
  cancel: the feeder blocks on `jobs <- i`, the workers have exited, `wg.Wait()`
  hangs.
- Closing `jobs` from a worker, or from several places, panics with "close of
  closed channel". Exactly one goroutine — the feeder — owns that close.
- Writing `firstErr` without the mutex is a data race, and `go test -race` will
  say so even though the value happens to look right.
- Returning partial results alongside the error invites a caller to publish a
  half-rendered batch. Return `nil` with the error.
