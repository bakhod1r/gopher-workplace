# Protecting a Shared Counter

## Intuition

`healthy++` compiles to load, add, store. Two goroutines can both load 7, both
store 8, and one probe is lost. A mutex makes the three steps indivisible with
respect to other goroutines, which is exactly what the counter needs.

## Approach

1. Declare `wg`, `mu`, and `healthy` in the enclosing function.
2. Per host, start a goroutine that calls `check(host)` with no lock held.
3. If healthy, `mu.Lock()`, increment, `mu.Unlock()`; then `wg.Wait()` and return the count.

## Solution

```go
import "sync"

func CountHealthy(hosts []string, check func(string) bool) int {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		healthy int
	)

	for _, host := range hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			if !check(host) {
				return
			}

			mu.Lock()
			healthy++
			mu.Unlock()
		}(host)
	}

	wg.Wait()
	return healthy
}
```

## Walkthrough

For a fleet where three of five hosts answer, the three healthy goroutines
queue on the mutex and increment one at a time, giving 3. The two unhealthy
ones never take the lock at all.

## Pitfalls

- Incrementing without the mutex — `go test -race` reports it and the count comes out low.
- Holding the lock across `check`, which serialises the probes and makes the whole round as slow as a sequential one.
- Reading `healthy` before `wg.Wait()` returns, which reports a partial count.
