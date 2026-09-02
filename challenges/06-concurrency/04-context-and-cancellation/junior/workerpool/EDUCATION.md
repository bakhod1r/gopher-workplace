# Stopping a Worker Pool

## Intuition

One `cancel()` closes one channel, and every goroutine blocked on it wakes: that is cancellation as a broadcast, and it is why a pool of any size needs no per-worker signalling. Pre-sizing the result slice means each worker writes a distinct element, so there is no shared memory to protect — appending instead would be a genuine data race.

## Approach

1. Create the pool context; `defer cancel()`.
2. `reasons := make([]error, n)`.
3. For each `i`: `wg.Add(1)`, start a goroutine that waits on `ctx.Done()` and writes `reasons[i] = ctx.Err()`.
4. Call `cancel()`, then `wg.Wait()`, then return `reasons`.

## Solution

```go
// StopWorkers starts n image-resize workers that are idle, waiting on the pool
// context, then triggers the pool shutdown and waits for every worker to
// finish. It returns the stop reason each worker observed, indexed by worker
// number.
//
// Every element is context.Canceled, and the slice has exactly n elements.
//
// Examples:
//
//	StopWorkers(0)  => []
//	StopWorkers(1)  => [context.Canceled]
//	StopWorkers(3)  => [context.Canceled, context.Canceled, context.Canceled]
func StopWorkers(n int) []error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reasons := make([]error, n)

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-ctx.Done()
			reasons[i] = ctx.Err()
		}(i)
	}

	cancel()
	wg.Wait()

	return reasons
}
```

## Walkthrough

- All `n` workers park on the same `Done()` channel.
- `cancel()` closes it once, and every worker's receive returns.
- Each worker stores `context.Canceled` at its own index and calls `wg.Done()`; `wg.Wait()` returns only when the last one has finished, so the slice is fully written before it is returned.

## Pitfalls

- `reasons = append(reasons, err)` from multiple goroutines is a data race — `-race` will fail the build.
- `wg.Add(1)` must happen before the goroutine is started, not inside it.
- Returning without `wg.Wait()` can hand back a partly written slice and leaks the workers.
- With `n == 0` the function must still return a non-nil empty slice of length 0.
