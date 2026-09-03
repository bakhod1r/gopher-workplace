# First Shard Error Cancels the Rest

## Intuition

`errgroup.WithContext` bundles the two things a fan-out always needs: waiting for everybody, and agreeing to quit early. The group records the *first* non-nil error and cancels its context at the same instant, so "first error wins" and "everyone else stops" are one event, not two you have to coordinate by hand.

## Approach

1. `group, groupCtx := errgroup.WithContext(ctx)`.
2. For each shard, `group.Go(func() error { return shard(groupCtx) })`.
3. `return group.Wait()`.

## Solution

```go
import (
	"context"

	"golang.org/x/sync/errgroup"
)
// ShardQuery runs the analytics query against a single shard.
type ShardQuery func(ctx context.Context) error

// QueryAllShards runs one query per shard concurrently and waits for all of
// them. Every shard receives the errgroup's context, so the first shard that
// fails cancels the others and the dashboard request stops paying for scans
// whose results are already useless.
//
// It returns nil when every shard succeeded, otherwise the first error the
// group recorded.
//
// Examples:
//
//	QueryAllShards(ctx, nil)                    => nil
//	QueryAllShards(ctx, [ok, ok])               => nil
//	QueryAllShards(ctx, [broken, parked, parked]) => the broken shard's error
func QueryAllShards(ctx context.Context, shards []ShardQuery) error {
	group, groupCtx := errgroup.WithContext(ctx)
	for _, shard := range shards {
		group.Go(func() error {
			return shard(groupCtx)
		})
	}
	return group.Wait()
}
```

## Walkthrough

- With one broken shard and two parked ones, the broken shard returns immediately; the group stores that error and cancels `groupCtx`; both parked shards wake, observe `context.Canceled`, and return. `Wait` reports the broken shard's error.
- The `later_error_is_ignored` case proves the group keeps the first error: the late shard cannot even run until cancellation has already happened.
- With an already-cancelled parent, `groupCtx` is born cancelled, so the parked shard returns `context.Canceled` straight away.
- An empty shard list makes `Wait` return nil with no goroutines started.

## Pitfalls

- Capturing and passing the outer `ctx` into `group.Go` — the shards then never learn about the first failure.
- Ignoring `group.Wait()`'s return value, which both drops the error and skips waiting.
- Assuming `Wait` aggregates every error: it returns exactly one, the first recorded.
- Calling `group.Go` after `Wait` has returned — the group is finished, and the new goroutine is unsupervised.
