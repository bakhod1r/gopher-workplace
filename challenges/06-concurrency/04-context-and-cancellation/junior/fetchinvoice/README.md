# Context as the First Parameter

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The billing store's read path is being brought in line with the rest of the codebase: every function that can block or do I/O takes a `context.Context` as its first parameter, named `ctx`, and honours it before doing work. The lookup also has its own domain error for a missing invoice, and the two must not be confused.

## Task

Implement the exported function(s) in [fetchinvoice.go](fetchinvoice.go) so that:

1. `ctx` is the first parameter and is checked first: a finished context returns `"", ctx.Err()`.
2. A non-positive ID returns `"", ErrNotFound`.
3. Otherwise it returns `fmt.Sprintf("invoice-%d", id)` and `nil`.
4. The context check takes precedence over the ID check.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  FetchInvoice(live ctx, 7)
Output: "invoice-7", nil
```

**Example 2:**

```
Input:  FetchInvoice(live ctx, 0)
Output: "", ErrNotFound
```

**Example 3:**

```
Input:  FetchInvoice(cancelled ctx, 7)
Output: "", context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **ctx first, named `ctx`** | The convention every stdlib and third-party Go API follows. |
| 2 | **Context error vs domain error** | "Nobody is waiting" is different from "the record does not exist". |
| 3 | **Sentinel errors** | `ErrNotFound` lets callers branch with `errors.Is`. |

## Hint

Check `ctx.Err()` first, then validate `id`, then build the result.

## Validate

```bash
make verify
```
