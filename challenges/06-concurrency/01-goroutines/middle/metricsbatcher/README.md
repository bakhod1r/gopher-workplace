# Metrics Batch Flush

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

The telemetry agent buffers metric points and ships them to the collector in batches. Batches go out in parallel because the collector is remote and slow, but a rejected batch is retried as a unit on the next tick — and the retry buffer has to preserve the original point order or the collector's own de-duplication window breaks.

## Task

Implement the exported function(s) in [metricsbatcher.go](metricsbatcher.go) so that:

1. Cut `points` into consecutive batches of at most `batchSize` points; the last batch may be short.
2. Flush each batch in its own goroutine, joined with a `sync.WaitGroup`.
3. Count the points of every accepted batch.
4. Re-queue the points of every rejected batch, in original input order.
5. With `batchSize <= 0` call `flush` zero times and re-queue every point.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  FlushBatches([]int{1, 2, 3, 4}, 2, flush)
Output: 4, []
```

**Example 2:**

```
Input:  FlushBatches([]int{-1, 2, 3, -4, 5}, 2, flush)
Output: 1, [-1 2 3 -4]
```

**Example 3:**

```
Input:  FlushBatches(nil, 2, flush)
Output: 0, []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Chunking then fanning out** | Build the batch boundaries first; the goroutines then own disjoint sub-slices. |
| 2 | **Sub-slice aliasing** | `points[start:end]` shares the backing array — safe here because each goroutine only reads. |
| 3 | **Order-preserving flatten** | Walking the batches in index order after `Wait` restores the input order for free. |
| 4 | **Degenerate input** | A non-positive batch size is a caller bug; fail closed by re-queueing everything. |

## Hint

Materialise `[][]int` first so each goroutine gets a stable batch index, then fold the accepted counts and the retry buffer in the parent, in order.

## Validate

```bash
make verify
```
