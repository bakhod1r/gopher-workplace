# Nightly Batch Cut-off

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The settlement batch must finish before the clearing house closes at 04:00. The scheduler hands the job an absolute cut-off instant, not a duration. If an earlier job overran and this one starts after the cut-off, it must refuse immediately and report a deadline failure so the on-call runbook routes it to the "missed window" playbook.

## Task

Implement the exported function(s) in [batchcutoff.go](batchcutoff.go) so that:

1. It derives a context from `context.Background()` with the deadline `cutoff`.
2. It waits on `ctx.Done()` and returns `ctx.Err()`.
3. For any instant already past, the result is `context.DeadlineExceeded`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  MissedCutoff(time.Now().Add(-time.Hour))
Output: context.DeadlineExceeded
```

**Example 2:**

```
Input:  MissedCutoff(time.Unix(0, 0))
Output: context.DeadlineExceeded
```

**Example 3:**

```
Input:  MissedCutoff(time.Time{})
Output: context.DeadlineExceeded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.WithDeadline`** | Takes an absolute `time.Time` instead of a duration. |
| 2 | **Deadline vs timeout** | `WithTimeout(p, d)` is exactly `WithDeadline(p, time.Now().Add(d))`. |
| 3 | **`context.DeadlineExceeded`** | Reported whether the deadline passed while waiting or was already past. |

## Hint

`context.WithDeadline(context.Background(), cutoff)`, `defer cancel()`, `<-ctx.Done()`, return `ctx.Err()`.

## Validate

```bash
make verify
```
