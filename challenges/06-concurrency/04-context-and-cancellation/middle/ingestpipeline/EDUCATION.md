# Cancel Propagates Through the Pipeline

## Intuition

A pipeline is only as cancellable as its most blocked stage. The consumer can `return` the instant a record is rejected, but the producer is sitting in `raw <- record` and will sit there forever unless something closes the door behind it. That something is `pipeCtx`: cancelling it makes the producer's `select` take the other arm, `close(raw)` runs, and the run winds down with no leaked goroutine.

## Approach

1. Return early if `ctx` is already finished.
2. `pipeCtx, cancel := context.WithCancel(ctx)`; `defer cancel()`.
3. Producer goroutine: `defer close(raw)`, then for each record `select` on `raw <- record` and `<-pipeCtx.Done()`.
4. Consumer: `for record := range raw`, call `normalise(pipeCtx, record)`, `cancel()` and return `nil, err` on failure.
5. After a clean drain, return the collected slice.

## Solution

```go
import (
	"context"
)
// Normalise cleans one raw event record. Returning an error aborts the whole
// ingest run.
type Normalise func(ctx context.Context, record string) (string, error)

// RunIngestPipeline streams raw event records through a producer goroutine into
// a normalising stage and collects the results. Producer and consumer share one
// derived context: whether the run ends from a bad record, a cancelled upload
// or an expired budget, the same cancel unblocks the producer so it can never
// be left parked on a send nobody will receive.
//
// It returns the normalised records in input order, or nil and the reason the
// run was abandoned.
//
// Examples:
//
//	RunIngestPipeline(ctx, ["a" "b"], upper)          => ["A" "B"], nil
//	RunIngestPipeline(ctx, ["a" "bad" "c"], upper)    => nil, errBadRecord
//	RunIngestPipeline(cancelled ctx, ["a"], upper)    => nil, context.Canceled
func RunIngestPipeline(ctx context.Context, records []string, normalise Normalise) ([]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	pipeCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	raw := make(chan string)
	go func() {
		defer close(raw)
		for _, record := range records {
			select {
			case raw <- record:
			case <-pipeCtx.Done():
				return
			}
		}
	}()

	clean := make([]string, 0, len(records))
	for record := range raw {
		value, err := normalise(pipeCtx, record)
		if err != nil {
			cancel()
			return nil, err
		}
		clean = append(clean, value)
	}
	if err := pipeCtx.Err(); err != nil {
		return nil, err
	}
	return clean, nil
}
```

## Walkthrough

- Three good records flow one at a time through the unbuffered channel, so the output order matches the input order without any sorting.
- `["alpha", "bad", "gamma"]` calls `normalise` exactly twice: the rejection returns before `gamma` is ever requested, and the `cancel()` releases the producer parked on its third send.
- An already-cancelled upload returns before the goroutine starts, so the normaliser is never called at all.
- The final `pipeCtx.Err()` check covers the case where the producer exited through the `Done()` arm and closed the channel: the drain looked clean, but the run was actually abandoned.

## Pitfalls

- A bare `raw <- record` with no `select`: the producer deadlocks the moment the consumer breaks out early.
- Closing `raw` from the consumer — a second close, or a send on a closed channel, panics.
- Passing the parent `ctx` to `normalise` instead of `pipeCtx`, so a pipeline-internal abort is invisible to the stage.
- Returning the partially collected slice alongside an error: the caller cannot tell a complete batch from a truncated one.
