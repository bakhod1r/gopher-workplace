# Metrics Batch Flush

## Intuition

The chunking is what makes this a clean fan-out: after the boundaries exist, batch `i` is a private, read-only window into the input, and goroutine `i` can neither see nor disturb its neighbours. All the shared state is deferred to a sequential fold at the end.

## Approach

1. Handle `batchSize <= 0` by returning `0` and a copy of every point.
2. Slice `points` into `batches [][]int` using a `start += batchSize` loop, clamping the final `end`.
3. Launch one goroutine per batch writing `flush(batch) != nil` into `failed[i]`.
4. `wg.Wait()`, then walk the batches in order: add `len(batch)` to the total or append the batch to the retry buffer.

## Solution

```go
// Package metricsbatcher — Gopher Workplace challenge.
package metricsbatcher

import "sync"

// FlushBatches cuts points into batches of batchSize, flushes each batch in its
// own goroutine, and reports how many points the collector accepted plus the
// points that must be re-queued, in their original order. A rejected batch is
// retried whole, so its points come back untouched.
//
// A batchSize of zero or less flushes nothing and re-queues everything.
//
// Examples:
//
//	FlushBatches([]int{1, 2, 3, 4}, 2, flush)  => 4, []
//	FlushBatches([]int{1, -1, 3}, 2, flush)    => 1, [1 -1]
//	FlushBatches(nil, 2, flush)                => 0, []
func FlushBatches(points []int, batchSize int, flush func(batch []int) error) (int, []int) {
	if batchSize <= 0 {
		retry := make([]int, len(points))
		copy(retry, points)
		return 0, retry
	}

	var batches [][]int
	for start := 0; start < len(points); start += batchSize {
		end := start + batchSize
		if end > len(points) {
			end = len(points)
		}
		batches = append(batches, points[start:end])
	}

	failed := make([]bool, len(batches))
	var wg sync.WaitGroup
	for i, batch := range batches {
		wg.Add(1)
		go func(i int, batch []int) {
			defer wg.Done()
			failed[i] = flush(batch) != nil
		}(i, batch)
	}
	wg.Wait()

	accepted := 0
	retry := []int{}
	for i, batch := range batches {
		if failed[i] {
			retry = append(retry, batch...)
			continue
		}
		accepted += len(batch)
	}
	return accepted, retry
}
```

## Walkthrough

- `[1 2 3 4]` at size 2 becomes two batches, both accepted, so all four points count.
- In `retry_keeps_order` batches `[-1 2]` and `[3 -4]` are rejected and `[5]` succeeds; the fold appends the rejected batches in index order, giving `[-1 2 3 -4]`.
- `ragged_last_batch` produces batches of 2, 2, and 1 — the clamp on `end` is what prevents a slice bounds panic.
- With `batchSize == 0` the guard returns before any goroutine starts, so `flush` is never called.

## Pitfalls

- Appending to the retry buffer from inside the goroutines — the point order then follows the scheduler.
- Forgetting to clamp `end` to `len(points)`, which panics on a ragged final batch.
- Counting points of a rejected batch as accepted because the error was checked after the increment.
- Returning the caller's slice as the retry buffer when `batchSize <= 0` instead of a copy.
