# Peek Job

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The job runner's `/status` HTTP handler peeks at the pending-job queue. The
handler must never block a request, so if no job is queued it reports the
runner as idle immediately.

## Task

Implement `PeekJob` in [queuepeek.go](queuepeek.go) so that:

1. It uses `select` with a `default` case.
2. It returns the job id and `true` when a receive can proceed right away.
3. It returns `0, false` when the queue is open and empty — it must not block.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PeekJob(5)
Output: 5, true
```

**Example 2:**

```
Input:  PeekJob() // open, empty
Output: 0, false
```

**Example 3:**

```
Input:  PeekJob() // closed
Output: 0, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`select` + `default`** | Makes a channel operation non-blocking. |
| 2 | **Readiness** | A closed channel is *always* ready — it yields the zero value. |
| 3 | **Non-blocking peek** | Required when an HTTP handler must never wait. |

## Hint

`default` runs only when no other case is ready. Note a closed channel
**is** ready, so it takes the receive case.

## Validate

```bash
make verify
```
