# Worker Job Queue

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A background worker pool pulls jobs from an in-memory queue. When the queue is empty a worker must block until a job arrives, not spin on the CPU. A `sync.Cond` parks the worker and wakes it when a producer enqueues.

## Task

Implement the stubbed functions in [jobqueue.go](jobqueue.go) so that:

1. `Submit` appends a job and wakes one waiting worker.
2. `Take` blocks while the queue is empty and open, then returns the oldest job; after `Close` with an empty queue it returns `"", false`.
3. `Close` marks the queue closed and wakes every waiting worker.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  q := NewJobQueue(); q.Submit("a"); q.Take()
Output: "a", true
```

**Example 2:**

```
Input:  q.Submit("a"); q.Submit("b"); q.Take(); q.Take()
Output: "a", true then "b", true
```

**Example 3:**

```
Input:  q := NewJobQueue(); q.Close(); q.Take()
Output: "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Cond** | `Wait` atomically unlocks the mutex, parks the goroutine, and re-locks on wake. |
| 2 | **for, not if** | Re-check the condition in a `for` loop: a wake-up does not guarantee the condition holds. |
| 3 | **Signal vs Broadcast** | One new job wakes one worker; closing wakes them all. |

## Hint

`for len(q.jobs) == 0 && !q.closed { q.notEmpty.Wait() }` — always a `for`, never an `if`.

## Validate

```bash
make verify
go test -race ./...
```
