# Shut Services Down in Reverse Order

## Intuition

Shutdown is startup played backwards. Every service that was started depended on the ones started before it, so tearing down in reverse is the only order in which nothing is pulled out from under a live dependant. The drain context is what makes it bounded: re-checking it *between* services means a closed window stops the sequence at a clean boundary rather than halfway through a single `Stop`.

## Approach

1. `stopped := make([]string, 0, len(services))` — non-nil from the start.
2. Loop `i` from `len(services)-1` down to 0.
3. Return `stopped, ctx.Err()` if the drain window has finished.
4. Return `stopped, err` if `services[i].Stop(ctx)` failed.
5. Append `services[i].Name` and continue; return `stopped, nil` at the end.

## Solution

```go
import (
	"context"
)
// Service is one subsystem of the API server, in startup order.
type Service struct {
	Name string
	Stop func(ctx context.Context) error
}

// ShutdownServices stops services in reverse startup order — the HTTP listener
// before the cache, the cache before the database — so that nothing is torn
// down while something above it still depends on it. Every Stop shares the
// drain context, and the sequence aborts as soon as the drain window closes or
// a Stop reports a failure.
//
// It returns the names of the services that stopped cleanly, in the order they
// were stopped, along with the reason the sequence ended (nil if it completed).
//
// Examples:
//
//	ShutdownServices(ctx, [db, cache, http])     => ["http" "cache" "db"], nil
//	ShutdownServices(ctx, [db, brokenCache, http]) => ["http"], errStopFailed
//	ShutdownServices(cancelled ctx, [db, http])  => [], context.Canceled
func ShutdownServices(ctx context.Context, services []Service) ([]string, error) {
	stopped := make([]string, 0, len(services))
	for i := len(services) - 1; i >= 0; i-- {
		if err := ctx.Err(); err != nil {
			return stopped, err
		}
		if err := services[i].Stop(ctx); err != nil {
			return stopped, err
		}
		stopped = append(stopped, services[i].Name)
	}
	return stopped, nil
}
```

## Walkthrough

- `[database, cache, http]` is walked as http → cache → database, so the returned names read `http,cache,database`.
- When the cache is stuck, the HTTP listener has already stopped; the function returns `["http"]` with the stop error, and the database is deliberately left running so the operator can investigate.
- If the very first stop (the HTTP listener) fails, nothing was stopped and the slice is empty but non-nil.
- A SIGTERM that has already cancelled the context, or an exhausted drain window, ends the sequence before the first `Stop` is called.

## Pitfalls

- Iterating forwards: the database goes away while the HTTP listener is still serving requests.
- Returning `nil, err` on failure and throwing away the partial progress the operator needs.
- Checking the context once before the loop instead of before each service.
- Continuing past a failed `Stop`: the failure means the layer below may still be in use.
