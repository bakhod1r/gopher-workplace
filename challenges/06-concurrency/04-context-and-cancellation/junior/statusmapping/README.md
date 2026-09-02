# Mapping Context Errors to Status

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The gateway's access log and its SLO dashboard need one label per request. Client disconnects (nginx logs them as 499) must not count against the availability SLO, while genuine timeouts must. The error bubbling up from the handler has usually been wrapped once or twice with `%w` on the way, so the classifier has to see through the wrapping.

## Task

Implement the exported function(s) in [statusmapping.go](statusmapping.go) so that:

1. `nil` maps to `"ok"`.
2. Anything that is or wraps `context.Canceled` maps to `"client_closed_request"`.
3. Anything that is or wraps `context.DeadlineExceeded` maps to `"gateway_timeout"`.
4. Everything else maps to `"internal_error"`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Status(nil)
Output: "ok"
```

**Example 2:**

```
Input:  Status(context.Canceled)
Output: "client_closed_request"
```

**Example 3:**

```
Input:  Status(fmt.Errorf("query: %w", context.DeadlineExceeded))
Output: "gateway_timeout"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`errors.Is`** | Unwraps `%w` chains and compares against a sentinel. |
| 2 | **Context sentinels** | `Canceled` and `DeadlineExceeded` mean very different things to an SLO. |
| 3 | **Tagless `switch`** | `switch { case cond: }` reads better than nested ifs for a classifier. |

## Hint

A `switch` with no tag: check `err == nil` first, then `errors.Is` for each sentinel.

## Validate

```bash
make verify
```
