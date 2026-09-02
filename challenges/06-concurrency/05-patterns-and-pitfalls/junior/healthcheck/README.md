# Health Check Counter

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The load balancer probes its whole backend fleet in parallel every few seconds
and needs one number: how many backends answered. `healthy++` looks atomic in
source but is a read, an add, and a write — run it from fifty goroutines and
probes go missing.

## Task

Implement `CountHealthy` in [healthcheck.go](healthcheck.go) so that:

1. It starts one goroutine per host, tracked by a `WaitGroup`.
2. A goroutine whose host fails the check returns without touching the counter.
3. A healthy host takes the mutex, increments the counter, and unlocks; the count is returned after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountHealthy([]string{"ok-a", "bad"}, isOK)
Output: 1
```

**Example 2:**

```
Input:  CountHealthy([]string{"ok-a", "ok-b"}, isOK)
Output: 2
```

**Example 3:**

```
Input:  CountHealthy(nil, isOK)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mutex** | Serialising a read-modify-write so no update is lost. |
| 2 | **Critical section size** | Lock around the increment only — never around `check`. |
| 3 | **WaitGroup** | `Wait` is what makes the final read of the counter safe. |

## Hint

Call `check` *outside* the lock and return early when it is false; only the
`healthy++` belongs in the critical section.

## Validate

```bash
make verify
```
