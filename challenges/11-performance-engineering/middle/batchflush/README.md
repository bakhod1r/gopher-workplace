# Amortising A Fixed Cost

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

One insert per row, one syscall per byte, one round trip per message: the per-item work is trivial and the per-*call* cost dominates. Batching divides that fixed cost by the batch size, and the only tricky parts are the partial batch at the end and the buffer that must not be reallocated on every flush.

## Task

Implement the three methods in [batchflush.go](batchflush.go):

1. `Add` appends an item and calls `Flush` when the batch reaches `Size`; a non-positive `Size` batches one item.
2. `Close` flushes a partial batch, is safe to call repeatedly, and never calls `Flush` with an empty batch.
3. A nil `Flush` must not panic but still counts as a flush; steady-state batching must not allocate.

## Examples

**Example 1:**

```
Input:  Size 2; Add(1..4)
Output: batches [1 2] and [3 4]
```

**Example 2:**

```
Input:  Size 3; Add(1); Add(2); Close()
Output: one batch [1 2]
```

**Example 3:**

```
Input:  Size 2; Add(1); Add(2); Close(); Close()
Output: exactly one flush
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The tail batch** | Whatever is buffered when the stream ends still has to go somewhere. |
| 2 | **Reset, do not reallocate** | `buf[:0]` after each flush is what keeps the steady state allocation-free. |
| 3 | **Batching trades latency for throughput** | An item may wait for the batch to fill; that is the deal you are making. |

## Topics used again

Function fields, slice reuse, methods on pointer receivers, guards.

## Hint

One unexported helper that flushes and resets, called from both `Add` and `Close`.

## Validate

```bash
make verify
```
