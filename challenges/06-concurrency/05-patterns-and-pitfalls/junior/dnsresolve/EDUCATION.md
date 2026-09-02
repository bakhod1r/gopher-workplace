# Guarding a Shared Map

## Intuition

Go deliberately does not make maps concurrency-safe, and it does not fail
quietly either: the runtime detects concurrent writes and kills the program.
A mutex around the write is the direct fix; keeping the lookup outside it is
what preserves the speed-up.

## Approach

1. `addrs := make(map[string]string, len(hosts))` plus a `WaitGroup` and a `sync.Mutex`.
2. Per host, a goroutine calls `lookup(host)` unlocked.
3. Lock, assign `addrs[host]`, unlock; `wg.Wait()` then return the map.

## Solution

```go
import "sync"

func ResolveAll(hosts []string, lookup func(string) string) map[string]string {
	addrs := make(map[string]string, len(hosts))

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	for _, host := range hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()

			addr := lookup(host)

			mu.Lock()
			addrs[host] = addr
			mu.Unlock()
		}(host)
	}

	wg.Wait()
	return addrs
}
```

## Walkthrough

Five hosts resolve at the same time. Each goroutine finishes its lookup, then
queues briefly for the mutex to store one entry. The lock is held for a single
assignment, so the lookups overlap almost completely.

## Pitfalls

- Writing `addrs[host] = lookup(host)` with no lock: "fatal error: concurrent map writes" in production, intermittently.
- Holding the mutex across the lookup, which serialises every resolution.
- Returning the map before `wg.Wait()`, exposing a half-filled map to the caller.
