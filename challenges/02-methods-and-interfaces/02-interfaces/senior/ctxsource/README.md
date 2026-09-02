# Context-Aware Source

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A long scan must abandon its work the moment the caller's context is cancelled, rather than running to completion.

## Task

Implement the stub(s) in [ctxsource.go](ctxsource.go):

1. Implement `Next` on `*RangeSource`.
2. Implement `SumWithContext`, which sums a source but stops and returns `ctx.Err()` when the context is done.
3. Constraint: check cancellation every iteration — a cancelled scan of a 10M-element source must return promptly, not after the full scan.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a live context over [1 2 3]
Output: 6, nil
```

**Example 2:**

```
Input:  a context cancelled up front
Output: 0, context.Canceled
```

**Example 3:**

```
Input:  cancellation mid-scan
Output: partial sum discarded, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **context cancellation** | `ctx.Done()` is the standard stop signal. |
| 2 | **select with default** | A non-blocking cancellation check inside a tight loop. |
| 3 | **Streaming interfaces** | Reused: the source is drained one element at a time. |

## Hint

`select { case <-ctx.Done(): return 0, ctx.Err(); default: }` — the `default` keeps it non-blocking.

## Validate

```bash
make verify
```
