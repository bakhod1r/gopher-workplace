# First Row or Disconnect

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A search endpoint streams rows from the database over a channel while the HTTP handler waits for the first one to start writing the response. If the browser disconnects or the request budget runs out, waiting for a row that may never arrive would pin the handler goroutine forever — so the wait must watch the context as well.

## Task

Implement the exported function(s) in [firstrow.go](firstrow.go) so that:

1. It selects on `ctx.Done()` and on a receive from `rows`.
2. On a row, it returns the row and a `nil` error.
3. On context completion, it returns `""` and `ctx.Err()`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  FirstRow(live ctx, chan holding "alice")
Output: "alice", nil
```

**Example 2:**

```
Input:  FirstRow(cancelled ctx, empty chan)
Output: "", context.Canceled
```

**Example 3:**

```
Input:  FirstRow(expired ctx, empty chan)
Output: "", context.DeadlineExceeded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`select` on two channels** | Whichever is ready first wins; if both are ready the choice is random. |
| 2 | **`ctx.Done()` as a case** | The standard way to make any blocking receive cancellable. |
| 3 | **Returning `ctx.Err()`** | Propagates the exact reason to the caller. |

## Hint

A two-case `select`: `case <-ctx.Done():` returns `"", ctx.Err()`; `case row := <-rows:` returns `row, nil`.

## Validate

```bash
make verify
```
