# Shutdown Once

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Several subsystems can each decide the service must stop: the signal handler,
the health checker, the config watcher. They all broadcast by closing the same
`quit` channel — and closing an already-closed channel panics, taking the
process down during the shutdown it was supposed to perform.

## Task

Implement `ShutdownOnce` in [shutdownonce.go](shutdownonce.go) so that:

1. It starts `closers` goroutines that each request shutdown concurrently.
2. `close(quit)` happens exactly once no matter how many requesters there are.
3. It waits for every requester before returning; with zero closers the channel stays open.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ShutdownOnce(quit, 1)
Output: quit closed, no panic
```

**Example 2:**

```
Input:  ShutdownOnce(quit, 10)
Output: quit closed exactly once, no panic
```

**Example 3:**

```
Input:  ShutdownOnce(quit, 0)
Output: quit left open
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once** | `once.Do(f)` runs `f` once and blocks other callers until it has finished. |
| 2 | **Close panics** | A second `close` on the same channel is a runtime panic, not a no-op. |
| 3 | **Idempotent shutdown** | Any component may request shutdown; only the first one performs it. |

## Hint

Declare the `sync.Once` outside the loop — one `Once` shared by every
goroutine is what makes the close happen a single time.

## Validate

```bash
make verify
```
