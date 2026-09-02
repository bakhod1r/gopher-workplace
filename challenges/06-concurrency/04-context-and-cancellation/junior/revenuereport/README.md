# Revenue Report Export

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The finance dashboard exports a revenue report by streaming millions of ledger rows from the warehouse. Users abandon those exports constantly — closing the tab, or hitting the export budget — and an export that keeps summing after nobody is listening holds a warehouse connection for minutes. The aggregation loop has to watch the context on every iteration.

## Task

Implement the exported function(s) in [revenuereport.go](revenuereport.go) so that:

1. It loops, selecting each iteration on `ctx.Done()` and a receive from `rows`.
2. It adds each amount to a running total.
3. When `rows` is closed it returns the total and `nil`.
4. When the context finishes it returns the partial total and `ctx.Err()`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  TotalRevenue(live ctx, closed chan 100,250,25)
Output: 375, nil
```

**Example 2:**

```
Input:  TotalRevenue(live ctx, closed empty chan)
Output: 0, nil
```

**Example 3:**

```
Input:  TotalRevenue(cancelled ctx, empty chan)
Output: 0, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`for { select { ... } }`** | The loop-with-cancellation shape for any streaming consumer. |
| 2 | **Comma-ok receive** | Detects the closed stream so the loop terminates. |
| 3 | **Partial results** | Returning what was accumulated *and* the error keeps both facts. |

## Hint

`for { select { case <-ctx.Done(): return total, ctx.Err(); case v, ok := <-rows: if !ok { return total, nil }; total += v } }`.

## Validate

```bash
make verify
```
