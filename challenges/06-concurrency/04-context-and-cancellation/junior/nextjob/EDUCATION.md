# Pulling the Next Job

## Intuition

Three outcomes share one wait: a job arrives, the queue is closed, or the worker is told to stop. A closed channel never blocks again — it returns the zero value immediately, forever — so the `ok` flag is the only way to tell "queue finished" from "job with an empty ID". Distinguishing the two errors lets the supervisor log a clean exit rather than an incident.

## Approach

1. `select` on `<-ctx.Done()` and `job, ok := <-jobs`.
2. Done → `"", ctx.Err()`.
3. `!ok` → `"", ErrQueueClosed`.
4. Otherwise → `job, nil`.

## Solution

```go
// NextJob takes the next job ID off the worker pool's queue. It stops early
// when the worker's context finishes during a rolling deploy, and reports
// ErrQueueClosed once the producer has closed the queue.
//
// Examples:
//
//	NextJob(live ctx, chan with "job-1")  => "job-1", nil
//	NextJob(live ctx, closed chan)        => "", ErrQueueClosed
//	NextJob(cancelled ctx, empty chan)    => "", context.Canceled
func NextJob(ctx context.Context, jobs <-chan string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case job, ok := <-jobs:
		if !ok {
			return "", ErrQueueClosed
		}
		return job, nil
	}
}
```

## Walkthrough

- A buffered queue makes the receive case ready, so the head job is returned.
- A closed queue also makes the receive ready, but `ok` is false, so `ErrQueueClosed` is returned.
- With a cancelled context and an empty open queue only the done case is ready, so `ctx.Err()` is returned.

## Pitfalls

- Dropping the `ok` check turns a closed queue into an infinite stream of empty job IDs at full CPU.
- `ErrQueueClosed` is a clean exit; do not fold it into `ctx.Err()` — the two need different handling upstream.
- Never assume an empty job ID means "no job"; only `ok == false` means that.
