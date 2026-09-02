# Retry Planner

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A queue consumer plans its retry schedule up front: for each pending job it
knows how many times the job has already failed, and the delay before the next
try doubles with every failure. The schedule is computed concurrently so
planning never blocks the consumer loop.

## Task

Implement `Backoffs` in [retryplanner.go](retryplanner.go) so that:

1. Return a slice of delays in milliseconds, the same length as `attempts`.
2. Delay `i` is `baseMs` doubled `attempts[i]` times; a non-positive attempt number yields `baseMs`.
3. Compute each delay in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Backoffs([]int{0, 1, 3}, 100)
Output: [100 200 800]
```

**Example 2:**

```
Input:  Backoffs([]int{-1}, 100)
Output: [100]
```

**Example 3:**

```
Input:  Backoffs(nil, 100)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Local accumulator** | `delay` is declared inside the goroutine, so each schedule is computed privately. |

## Hint

Start the accumulator at `baseMs` and double it in a loop. A loop that never
runs leaves `baseMs`, which handles `0` and negative attempts for free.

## Validate

```bash
make verify
```
