# Error Filter Stage

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Only a small fraction of parsed log records are worth paging on. The alerting
branch of the ingest pipeline sits behind a *filter* stage that drops
everything but error records, so the expensive indexing step sees far less
traffic.

## Task

Implement `ErrorFilter` in [errorfilter.go](errorfilter.go) so that:

1. It returns a new channel immediately.
2. A goroutine ranges over `records` and sends only those where `isError(rec)` is true.
3. The output channel is closed when `records` is drained, even if nothing was forwarded.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  records "ERR disk", "INFO ok", "ERR io"
Output: "ERR disk", "ERR io" then closed
```

**Example 2:**

```
Input:  records "INFO ok", "WARN slow"
Output: no records, closed
```

**Example 3:**

```
Input:  closed empty stream
Output: closed immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Filter stage** | A stage may emit fewer values than it consumes. |
| 2 | **Close on the empty path** | The output must close even when zero records pass the predicate. |
| 3 | **Predicate injection** | `isError` keeps the severity rules out of the plumbing. |

## Hint

Same shape as a transforming stage; just wrap the send in an `if`. Keep
`defer close(out)` outside the `if` so it always runs.

## Validate

```bash
make verify
```
