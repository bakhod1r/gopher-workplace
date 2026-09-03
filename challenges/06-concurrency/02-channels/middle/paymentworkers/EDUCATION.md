# Payment Capture Workers

## Intuition

Two unsynchronised things must line up: work going in and results coming out. Channels carry both, and the only tricky part is who closes what and when. Wait for the workers on a *separate* goroutine, so the caller stays free to drain results.

## Approach

1. Normalise `workers` to at least 1.
2. Create `jobs` and `results` channels.
3. Start `workers` goroutines ranging over `jobs`, sending a `Result` per charge.
4. Start a feeder that sends every charge then closes `jobs`.
5. Start a closer that runs `wg.Wait()` then `close(results)`.
6. Range over `results` on the caller's goroutine, filling the map.

## Solution

```go
import "sync"

func CaptureAll(charges []string, workers int, capture func(charge string) string) map[string]string {
	if workers < 1 {
		workers = 1
	}

	jobs := make(chan string)
	results := make(chan Result)

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for charge := range jobs {
				results <- Result{Charge: charge, Status: capture(charge)}
			}
		}()
	}

	go func() {
		for _, c := range charges {
			jobs <- c
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	out := make(map[string]string, len(charges))
	for r := range results {
		out[r.Charge] = r.Status
	}
	return out
}
```

## Walkthrough

With 200 charges and 8 workers: the feeder pushes IDs; whichever worker is free takes the next one, so at most 8 `capture` calls run at once. Each worker sends its `Result` and the caller inserts it into the map. When `jobs` closes, each worker's `range` ends and `wg.Done` fires; the closer's `wg.Wait()` returns, `results` closes, and the caller's loop ends.

## Pitfalls

- Calling `wg.Wait()` before draining `results` — the workers block on send and nothing progresses.
- Closing `results` from a worker: the others are still sending on it.
- Writing into the map from the workers instead of the collector — that is a data race.
