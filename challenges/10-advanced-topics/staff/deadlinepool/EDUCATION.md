# Stop When The Deadline Passes

## Intuition

A cancellable pool is not one that checks a flag between items — it is one where every point that can wait also watches the context. The join at the end is what turns "they will stop eventually" into a guarantee.

## Approach

1. Normalise `workers` and handle the empty input.
2. Start workers that `select` on the index channel and `ctx.Done()`.
3. Feed indices with a `select` that breaks out on cancellation.
4. Close the channel, `Wait`, then return `ctx.Err()` or the results.

## Solution

```go
import (
	"context"
	"sync"
)

// Process doubles every item using workers goroutines and returns the
// results in input order.
//
// If ctx is cancelled or its deadline passes first, Process returns the
// context's error, and every goroutine it started must have exited.
//
// Examples:
//
// 	Process(ctx, []int{1, 2}, 2) => []int{2, 4}, nil
func Process(ctx context.Context, items []int, workers int) ([]int, error) {
	if workers < 1 {
		workers = 1
	}
	out := make([]int, len(items))
	if len(items) == 0 {
		return out, ctx.Err()
	}
	if workers > len(items) {
		workers = len(items)
	}

	idx := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case i, ok := <-idx:
					if !ok {
						return
					}
					out[i] = items[i] * 2
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	feed:
	for i := range items {
		select {
		case idx <- i:
		case <-ctx.Done():
			break feed
		}
	}
	close(idx)
	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
```

## Walkthrough

On cancellation the feed loop breaks, the channel is closed, and each worker returns from whichever `select` it is in. `Wait` then returns and the error is reported.

## Pitfalls

- Returning on `ctx.Done()` without waiting, which leaves the workers running past the call.
- Feeding with a plain send, which blocks forever once the workers have exited.
- Returning partial results on cancellation when the caller expects all or nothing.
