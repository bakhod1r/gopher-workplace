# Measuring Bounded Concurrency

## Intuition

The semaphore enforces the ceiling; the counter observes what really happened
under it. The peak is only trustworthy if the increment and the comparison are
atomic together, which is why both live inside the same `Lock`/`Unlock` pair.

## Approach

1. `sem := make(chan struct{}, limit)`, plus `wg`, `mu`, `inFlight`, `peak`.
2. Each goroutine acquires a permit, then locks to increment and update the peak.
3. It calls `do(req)`, then locks again to decrement, and releases the permit; `wg.Wait()` and return `peak`.

## Solution

```go
import "sync"

func PeakInFlight(requests []string, limit int, do func(string)) int {
	sem := make(chan struct{}, limit)

	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		inFlight int
		peak     int
	)

	for _, req := range requests {
		wg.Add(1)
		go func(req string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			mu.Lock()
			inFlight++
			if inFlight > peak {
				peak = inFlight
			}
			mu.Unlock()

			do(req)

			mu.Lock()
			inFlight--
			mu.Unlock()
		}(req)
	}

	wg.Wait()
	return peak
}
```

## Walkthrough

With five requests and a limit of 2, at most two goroutines hold permits at any
instant, so `inFlight` never exceeds 2 and neither can `peak`. On a very fast
machine the first request may finish before the second starts, giving a peak of
1 — which is why the test asserts a range, not an exact value.

## Pitfalls

- Updating `peak` outside the lock: two goroutines can interleave the read and the write and record a bogus value.
- Decrementing before `do` returns, which under-reports the true concurrency.
- Asserting an exact peak in a test — scheduling is not deterministic, only the `limit` ceiling is.
