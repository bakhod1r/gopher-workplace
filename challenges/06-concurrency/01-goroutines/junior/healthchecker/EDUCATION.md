# Health Checker

## Intuition

A status page is a fan-out: N independent checks, one row each. Give every check
its own goroutine and its own output slot and there is no shared state left to
protect.

## Approach

1. Allocate `out := make([]bool, len(services))`.
2. Launch a goroutine per service, passing `i` and `service` as parameters.
3. Inside, call `probe` and store the healthy/unhealthy verdict at `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package healthchecker — Gopher Workplace challenge.
package healthchecker

import (
	"sync"
)

// CheckAll reports which services answered with a healthy status code.
//
// Examples:
//
//	CheckAll([]string{"api", "db"}, probe)  => [true false]
//	CheckAll([]string{"api"}, probe)        => [true]
//	CheckAll(nil, probe)                    => []
func CheckAll(services []string, probe func(service string) int) []bool {
	out := make([]bool, len(services))
	var wg sync.WaitGroup
	for i, service := range services {
		wg.Add(1)
		go func(i int, service string) {
			defer wg.Done()
			code := probe(service)
			out[i] = code >= 200 && code < 400
		}(i, service)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"api"` answers `200`, which is in range, so `out[0]` is `true`.
- `"db"` answers `503`, so `out[1]` is `false`.
- `"cache"` answers `301` — still under 400, so it counts as healthy.

## Pitfalls

- Probing on the parent goroutine and only using `go` for the assignment.
- Treating any non-`200` code as unhealthy; the range is `[200, 400)`.
- Forgetting `defer wg.Done()`, which deadlocks `Wait` forever.
