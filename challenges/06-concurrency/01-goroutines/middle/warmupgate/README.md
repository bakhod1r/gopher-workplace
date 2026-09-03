# Cache Warm-Up Gate

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

Before a new pod is added to the load balancer it warms a list of caches. The deploy gate wants one verdict, not a list — but it must still be the *same* verdict on every run, and every warm must be finished before the gate answers. Reporting the failure of the lowest-indexed key makes the message reproducible; waiting for all of them makes the pod's state knowable.

## Task

Implement the exported function(s) in [warmupgate.go](warmupgate.go) so that:

1. Warm each key in its own goroutine, joined with a `sync.WaitGroup`.
2. Attempt every key: an early failure must not skip the remaining warms.
3. Wait for all goroutines before returning, so none outlive the call.
4. Return the error belonging to the lowest-indexed failing key.
5. Return `nil` when every key warmed, and for an empty key list.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  WarmCaches([]string{"tags", "prices"}, warm)
Output: <nil>
```

**Example 2:**

```
Input:  WarmCaches([]string{"cold-a", "tags", "cold-z"}, warm)
Output: cold source: cold-a
```

**Example 3:**

```
Input:  WarmCaches(nil, warm)
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deterministic error selection** | "First" must mean first by index, never first by arrival. |
| 2 | **Goroutine lifetime** | `wg.Wait()` before returning guarantees no warm is still running after the gate answers. |
| 3 | **Per-index collection** | Collect everything, then decide — deciding inside the goroutines reintroduces the race. |
| 4 | **Error identity** | Returning the original error keeps `errors.Is` working for the caller. |

## Hint

Store every result in `make([]error, len(keys))`, then scan that slice in order after `wg.Wait()` and return the first non-nil entry.

## Validate

```bash
make verify
```
