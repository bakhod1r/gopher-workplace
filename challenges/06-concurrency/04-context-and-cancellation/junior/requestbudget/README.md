# Reading the Request Budget

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Before starting an upload to object storage, the client wants to know whether the caller imposed a cut-off. With a tight budget it sends one PUT and fails fast; with no budget it uses resumable multipart. The `context.Context` interface exposes exactly this query.

## Task

Implement the exported function(s) in [requestbudget.go](requestbudget.go) so that:

1. It returns `ctx.Deadline()` unchanged.
2. Contexts with no deadline yield the zero `time.Time` and `false`.
3. A context derived from one that has a deadline reports that same deadline.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Budget(context.Background())
Output: time.Time{}, false
```

**Example 2:**

```
Input:  Budget(ctx with deadline t)
Output: t, true
```

**Example 3:**

```
Input:  Budget(child of that ctx)
Output: t, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`ctx.Deadline()`** | One of the four `context.Context` methods; reports the cut-off and whether it exists. |
| 2 | **Deadline inheritance** | A derived context keeps the tightest deadline in its ancestry. |
| 3 | **Multiple return values** | Forwarding `(time.Time, bool)` needs no unpacking. |

## Hint

One line: `return ctx.Deadline()`.

## Validate

```bash
make verify
```
