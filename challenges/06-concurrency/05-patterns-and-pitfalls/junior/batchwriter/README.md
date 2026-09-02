# Batching Writer

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Writing rows to the warehouse one statement at a time is an order of
magnitude slower than batching them. The loader reads a row stream and groups
it into insert batches, always flushing whatever is left when the stream ends
so the tail of the file is not silently dropped.

## Task

Implement `BatchInserts` in [batchwriter.go](batchwriter.go) so that:

1. A size of zero or less drains the stream and returns nil.
2. Otherwise it accumulates rows and appends a full batch each time it reaches `size` rows.
3. When the stream closes, any partial batch is flushed as the final batch.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BatchInserts(stream a, b, c, size 2)
Output: [][]string{{"a", "b"}, {"c"}}
```

**Example 2:**

```
Input:  BatchInserts(stream a, b, size 2)
Output: [][]string{{"a", "b"}}
```

**Example 3:**

```
Input:  BatchInserts(stream a, size 0)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Batching** | Amortising a fixed per-operation cost over many stream items. |
| 2 | **Final flush** | The partial batch after the close is the bug most implementations miss. |
| 3 | **Fresh backing array** | Reallocate the batch after flushing so the appended slice is not overwritten. |

## Hint

After appending a full batch, start a *new* slice with `make` — reusing the
old one with `batch = batch[:0]` would mutate the batch you just stored.

## Validate

```bash
make verify
```
