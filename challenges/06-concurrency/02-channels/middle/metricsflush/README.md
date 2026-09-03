# Metric Batch Flusher

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

A metrics agent buffers samples in memory and flushes them to the remote write endpoint in batches. A batch leaves either when it is full or when the collector closes the input, and the flusher must never emit an empty batch or lose the final partial one.

## Task

Implement the stubbed function in [metricsflush.go](metricsflush.go) so that:

1. Emit a batch as soon as it holds `size` samples.
2. When `in` closes, emit the remaining partial batch — but never an empty one.
3. Close the output channel once the input is drained.
4. A `size` of zero or less means one sample per batch.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  in: 1,2,3 size 2
Output: [1 2] then [3]
```

**Example 2:**

```
Input:  in: 1,2 size 5
Output: [1 2]
```

**Example 3:**

```
Input:  in: (closed, empty) size 3
Output: no batches
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Producer owns the close | The goroutine that sends on `out` is the one that closes it — `defer close(out)` states that ownership. |
| 2 | `range` over a channel | The loop ends exactly when the input is closed and drained; that is the signal to flush the tail. |
| 3 | Fresh backing array | `batch = batch[:0]` would hand consumers a slice you keep writing into. Allocate a new one per batch. |
| 4 | Directional types | `<-chan int` in, `<-chan []int` out: the signature alone says who may send. |

## Hint

Accumulate into a slice; on reaching `size`, send it and allocate a **new** slice. After the `range` ends, send the tail if it is non-empty, then let `defer close(out)` fire.

## Validate

```bash
make verify
```
