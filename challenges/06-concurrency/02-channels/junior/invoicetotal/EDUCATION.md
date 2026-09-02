# Fan-Out Over One Channel

## Intuition

A channel is safe for many concurrent receivers, and each value is handed
to exactly one of them. That single fact gives you a work queue for free:
fill it, close it, and let N goroutines `range` over it. The split between
workers is nondeterministic, but the *total* is not — which is why the
billing figure is stable.

## Approach

1. Clamp `workers` to at least 1.
2. Buffer every line on `queue` and `close(queue)`.
3. Start `workers` goroutines; each sums what it receives and sends one partial.
4. Track them with a `sync.WaitGroup`; `wg.Wait()`, then `close(partials)`.
5. `range` over `partials` and return the total.

## Solution

```go
import "sync"

func TotalCents(lines []int, workers int) int {
	if workers < 1 {
		workers = 1
	}

	queue := make(chan int, len(lines))
	for _, cents := range lines {
		queue <- cents
	}
	close(queue)

	partials := make(chan int, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for cents := range queue {
				sum += cents
			}
			partials <- sum
		}()
	}
	wg.Wait()
	close(partials)

	total := 0
	for sum := range partials {
		total += sum
	}
	return total
}
```

## Walkthrough

With lines `100, 250, 99` and 2 workers: both range over the queue; maybe
one takes `100, 99` (partial 199) and the other `250`. The partials add up
to `449` whichever way the split lands.

## Pitfalls

- Closing `partials` before `wg.Wait()` makes a worker send on a closed channel — panic.
- Not closing the queue leaves every worker blocked in `range` and `wg.Wait()` never returns.
- Accumulating into one shared variable instead of per-worker partials is a data race; `-race` will catch it.
