# Fan-Out: Many Workers, One Channel

## Intuition

Multiple goroutines may receive from the same channel, and the runtime hands
each value to exactly one of them. That single fact is the whole fan-out
pattern — no work-stealing queue, no assignment logic of your own.

## Approach

1. Create `jobs` and a `results` channel buffered to `len(images)`.
2. Start `workers` goroutines, each `for img := range jobs { results <- render(img) }`, tracked by a `WaitGroup`.
3. Queue all images, `close(jobs)`, `wg.Wait()`, `close(results)`, drain, `sort.Strings`.

## Solution

```go
import (
	"sort"
	"sync"
)

func RenderThumbnails(images []string, workers int, render func(string) string) []string {
	jobs := make(chan string)
	results := make(chan string, len(images))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for img := range jobs {
				results <- render(img)
			}
		}()
	}

	for _, img := range images {
		jobs <- img
	}
	close(jobs)
	wg.Wait()
	close(results)

	var out []string
	for r := range results {
		out = append(out, r)
	}
	sort.Strings(out)
	return out
}
```

## Walkthrough

With three images and two renderers, renderer A might take `a` and `c` while B
takes `b` — or any other split. Results land in completion order, so sorting
is what makes the return value deterministic regardless of the split.

## Pitfalls

- Closing `results` before `wg.Wait()` — a renderer still running panics with "send on closed channel".
- Leaving `results` unbuffered and draining only after `wg.Wait()`: renderers block on send and `Wait` never returns.
- Expecting thumbnails back in upload order — completion order is not job order.
