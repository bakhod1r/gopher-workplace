# Cancel Propagates Through the Pipeline

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The event ingest endpoint streams an uploaded batch through a producer goroutine and a normalising stage. A rejected record must abandon the run — but abandoning it naively deadlocks: the producer is parked on a send into a channel the consumer has stopped reading. The derived pipeline context is the release valve, and it also carries the client's cancellation down into the normaliser.

## Task

Implement the exported function(s) in [ingestpipeline.go](ingestpipeline.go) so that:

1. It returns `nil, ctx.Err()` immediately if the ingest context is already finished, without starting a goroutine.
2. It derives `pipeCtx` with `context.WithCancel(ctx)` and `defer`s the cancel.
3. It starts a producer goroutine that sends each record with a `select` over the send and `<-pipeCtx.Done()`, and closes the channel on the way out.
4. The consumer normalises each record with `pipeCtx`; a normalise error cancels the pipeline and returns `nil` plus that error.
5. On success it returns the normalised records in input order; it must not leak the producer goroutine on any path.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  RunIngestPipeline(ctx, ["alpha" "beta"], upper)
Output: ["ALPHA" "BETA"], nil
```

**Example 2:**

```
Input:  RunIngestPipeline(ctx, ["alpha" "bad" "gamma"], upper)
Output: nil, errBadRecord  (normalise called twice)
```

**Example 3:**

```
Input:  RunIngestPipeline(cancelled ctx, ["alpha"], upper)
Output: nil, context.Canceled  (normalise never called)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cancel as a deadlock release** | A blocked sender needs a `Done()` arm or it parks forever once the consumer leaves. |
| 2 | **`defer close(chan)`** | The producer owns the channel and is the only one allowed to close it. |
| 3 | **Propagating one context through stages** | Producer and normaliser observe the same cancellation, so the whole pipeline stops together. |

## Hint

The producer's send must be a `select` with a `<-pipeCtx.Done()` arm — that is what lets the consumer walk away from a bad record without deadlocking the run.

## Validate

```bash
make verify
```
