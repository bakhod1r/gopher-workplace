# Startup Sequence

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A service runs a set of preflight checks before it accepts traffic: config
parses, disk has room, the network is reachable. The checks do not depend on one
another, so they all run at once, and the boot log prints their status codes in
the configured order rather than in the order they finished.

## Task

Implement `RunChecks` in [startupsequence.go](startupsequence.go) so that:

1. Run every function in `checks` in its own goroutine.
2. Store the return value of `checks[i]` at `out[i]` — the report order must match the input order.
3. Join the goroutines with a `sync.WaitGroup` before returning.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RunChecks([]func() int{configOK, diskFull})
Output: [0 28]
```

**Example 2:**

```
Input:  RunChecks([]func() int{configOK})
Output: [0]
```

**Example 3:**

```
Input:  RunChecks(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Functions are values** | A `func() int` is passed into a goroutine like any other value. |

## Hint

Position comes from the index `i`, not from which check finishes first. Pass
both `i` and `check` into the goroutine.

## Validate

```bash
make verify
```
