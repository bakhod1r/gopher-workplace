# Pulling the Next Job

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Thumbnail workers pull job IDs off an in-process queue fed by the API. During a rolling deploy each worker's context is cancelled and it must stop pulling; when the producer finishes a batch it closes the queue and workers must exit normally instead of spinning on a closed channel forever.

## Task

Implement the exported function(s) in [nextjob.go](nextjob.go) so that:

1. It selects on `ctx.Done()` and a receive from `jobs`.
2. On a job it returns the ID and `nil`.
3. On a closed channel (`ok == false`) it returns `""` and `ErrQueueClosed`.
4. On context completion it returns `""` and `ctx.Err()`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NextJob(live ctx, queue holding "job-1")
Output: "job-1", nil
```

**Example 2:**

```
Input:  NextJob(live ctx, closed queue)
Output: "", ErrQueueClosed
```

**Example 3:**

```
Input:  NextJob(cancelled ctx, empty queue)
Output: "", context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok receive** | `v, ok := <-ch` distinguishes a real value from a closed channel. |
| 2 | **Closed channels are always ready** | Without the `ok` check a worker spins on zero values forever. |
| 3 | **Sentinel errors** | `ErrQueueClosed` lets the caller exit cleanly instead of logging a failure. |

## Hint

Use `case job, ok := <-jobs:` and branch on `ok` before returning.

## Validate

```bash
make verify
```
