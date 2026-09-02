# A Cancelled Query Must Not Kill the Request

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The product page handler queries the recommendation service with its own derived context so it can abandon a slow call and fall back to a cached list. A reviewer worried that cancelling that call would take the whole request down with it and blank the page. Write the experiment that settles it.

## Task

Implement the exported function(s) in [childisolation.go](childisolation.go) so that:

1. It creates a cancellable request context and derives a query context from it.
2. It cancels only the query context and waits for its `Done()`.
3. It returns the *request* context's `Err()`, which is `nil`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  RequestErrAfterQueryCancel()
Output: nil
```

**Example 2:**

```
Input:  result == context.Canceled
Output: false
```

**Example 3:**

```
Input:  the request context stays usable
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **One-way propagation** | Cancellation flows from parent to child, never upward. |
| 2 | **Independent children** | Siblings are unaffected by each other's cancellation. |
| 3 | **Scoped abandonment** | A derived context is how you abandon one sub-operation safely. |

## Hint

Cancel the child, then read the *parent's* `Err()` — it is still nil.

## Validate

```bash
make verify
```
