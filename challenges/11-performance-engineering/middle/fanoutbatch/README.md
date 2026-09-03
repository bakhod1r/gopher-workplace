# Fan Out Over Batches, Not Items

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Launching a goroutine per item makes the scheduler the bottleneck: a goroutine costs a couple of kilobytes and a few hundred nanoseconds, which is more than the work when the item is small. Batching first means the fan-out width is the batch count, and the per-goroutine cost is amortised over a whole chunk.

## Task

Implement both functions in [fanoutbatch.go](fanoutbatch.go):

1. `Chunks` splits the items into contiguous batches of `size`, the last one possibly short, as sub-slices sharing the input's memory.
2. A non-positive `size` or no items gives an empty, non-nil result.
3. `SumBatches` sums each batch concurrently and returns the sums in batch order, race-free under `-race`.

## Examples

**Example 1:**

```
Input:  Chunks([1 2 3 4 5], 2)
Output: [[1 2] [3 4] [5]]
```

**Example 2:**

```
Input:  SumBatches([1 2 3 4 5], 2)
Output: [3 7 5]
```

**Example 3:**

```
Input:  Chunks([1 2], 0)
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutines are cheap, not free** | Around 2KB of stack and scheduling overhead — worth it per batch, not per item. |
| 2 | **Sub-slices share memory** | Chunking costs one slice header per batch, not a copy of the data. |
| 3 | **Order comes from the index** | Each goroutine writing `out[i]` keeps the results deterministic without sorting. |

## Topics used again

Slice windows, `min`, goroutines, `sync.WaitGroup`.

## Hint

`items[i:min(i+size, len(items))]` is the batch; its index is its slot in the result.

## Validate

```bash
make verify
```
