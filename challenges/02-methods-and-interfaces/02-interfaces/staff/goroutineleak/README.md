# Goroutine Leak

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Each request started a watcher goroutine. The watchers never exited, and after a week the process held hundreds of thousands of them.

## Task

Implement the stub(s) in [goroutineleak.go](goroutineleak.go):

1. Implement `Start` on `*Watcher` so the goroutine exits when either the input closes or `Stop` is called.
2. Implement `Stop`, which is idempotent and waits for the goroutine to finish.
3. Constraint: `-race` clean, and the test asserts the goroutine count returns to its baseline — a leaked goroutine fails the test.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Start then Stop
Output: the goroutine exits
```

**Example 2:**

```
Input:  Start, close the input
Output: the goroutine exits on its own
```

**Example 3:**

```
Input:  Stop twice
Output: safe
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutine lifetime** | Every goroutine needs an owner and a guaranteed exit path. |
| 2 | **select on a stop channel** | The standard shutdown signal alongside the work channel. |
| 3 | **Leak detection** | `runtime.NumGoroutine` turns a lifetime claim into an assertion. |

## Hint

`select` on both the input and `stop`; use `sync.Once` so `Stop` can be called repeatedly.

## Validate

```bash
make verify
```
