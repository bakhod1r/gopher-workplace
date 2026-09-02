# The Worker Pool

## Intuition

A pool decouples *how much work exists* from *how much runs at once*. Workers
block on `range jobs` when the queue is empty and exit when it closes, so the
pool tears itself down as soon as the batch ends.

## Approach

1. Create `jobs` (unbuffered) and `results` buffered to `len(uploads)`.
2. Start `workers` goroutines tracked by a `WaitGroup`, each resizing jobs into `results`.
3. Queue all uploads, close `jobs`, `wg.Wait()`, close `results`, drain, sort.

## Solution

```go
import (
	"sort"
	"sync"
)

func ResizeQueue(uploads []string, workers int, resize func(string) string) []string {
	jobs := make(chan string)
	results := make(chan string, len(uploads))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range jobs {
				results <- resize(key)
			}
		}()
	}

	for _, key := range uploads {
		jobs <- key
	}
	close(jobs)
	wg.Wait()
	close(results)

	out := make([]string, 0, len(uploads))
	for r := range results {
		out = append(out, r)
	}
	sort.Strings(out)
	if len(out) == 0 {
		return nil
	}
	return out
}
```

## Walkthrough

With 3 uploads and 2 workers: both workers grab a key, and the producer blocks
sending the third until one is free — that block *is* the backpressure.
`close(jobs)` ends both ranges, `wg.Wait()` returns, and the buffered results
hold all three resized keys in completion order.

## Pitfalls

- `wg.Wait()` before `close(jobs)`: the workers never see the close, so they never exit and `Wait` hangs.
- Starting one goroutine per upload instead of per worker — that is unbounded spawn, not a pool.
- Closing `results` from a worker: the other workers panic on their next send.
