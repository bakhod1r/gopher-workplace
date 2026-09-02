# Telemetry Flusher

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A telemetry agent buffers response codes and flushes them to the backend in
fixed-size batches. Each flush carries a summary of how many server errors
(status `500` and above) the batch contains, and the summaries are computed
concurrently so flushing never stalls the request path.

## Task

Implement `BatchErrorCounts` in [telemetryflusher.go](telemetryflusher.go) so that:

1. Return `nil` when `batch <= 0`.
2. Split `codes` into consecutive batches of `batch` entries; the last batch may be shorter.
3. Write the number of codes `>= 500` in batch `c` to `out[c]`, one goroutine per batch.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BatchErrorCounts([]int{200, 500, 503, 200}, 2)
Output: [1 1]
```

**Example 2:**

```
Input:  BatchErrorCounts([]int{500, 500, 200}, 2)
Output: [2 0]
```

**Example 3:**

```
Input:  BatchErrorCounts([]int{200}, 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Private counters** | Each goroutine counts into its own local variable; only the final value crosses out via `out[c]`. |

## Hint

Keep `errors` inside the goroutine. Incrementing one shared counter from several
goroutines is exactly the race this topic teaches you to avoid.

## Validate

```bash
make verify
```
