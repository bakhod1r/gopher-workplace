# Health Fan-Out

## Intuition

A fan-out that writes into one shared slice has exactly one shared resource: the slice header. Hold the lock for the nanosecond it takes to append, and the probes still run in parallel. Hold it across the probe and you have written a sequential loop with extra steps.

## Approach

1. Declare `mu sync.Mutex`, `down []string`, and `wg sync.WaitGroup`.
2. Launch a goroutine per service, passing the name in as a parameter.
3. Inside, `defer wg.Done()`; on a non-nil probe error lock, append, unlock.
4. `wg.Wait()`, `sort.Strings(down)`, and normalise a nil slice to `[]string{}`.

## Solution

```go
// Package healthfanout — Gopher Workplace challenge.
package healthfanout

import (
	"sort"
	"sync"
)

// UnhealthyServices probes every service concurrently and returns the names
// that failed their health check, sorted alphabetically so the on-call page
// reads the same way on every run.
//
// Examples:
//
//	UnhealthyServices([]string{"api", "db"}, probe)  => ["db"]
//	UnhealthyServices([]string{"api"}, probe)        => []
//	UnhealthyServices(nil, probe)                    => []
func UnhealthyServices(services []string, probe func(service string) error) []string {
	var (
		mu   sync.Mutex
		down []string
		wg   sync.WaitGroup
	)
	for _, svc := range services {
		wg.Add(1)
		go func(svc string) {
			defer wg.Done()
			if err := probe(svc); err != nil {
				mu.Lock()
				down = append(down, svc)
				mu.Unlock()
			}
		}(svc)
	}
	wg.Wait()
	sort.Strings(down)
	if down == nil {
		return []string{}
	}
	return down
}
```

## Walkthrough

- With three healthy services no goroutine ever appends, so the empty-slice normalisation is what the caller sees.
- In `all_down` both names are appended in arbitrary order and the sort turns them into `[api web]`.
- `result_is_sorted` starts from an unsorted input to prove the ordering comes from the sort, not from the input.
- The call count assertion proves every service is probed exactly once, regardless of outcome.

## Pitfalls

- Appending without the mutex — the race detector fails the build, and entries are silently lost.
- Locking around `probe(svc)` serialises the fan-out and defeats the point of the goroutines.
- Sorting before `wg.Wait()` sorts a partial result.
- Returning the nil slice directly, which marshals as `null` instead of `[]`.
