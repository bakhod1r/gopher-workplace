# Request ID for Logging

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Every log line the API emits should carry the request ID the edge proxy assigned, so an operator can pull one request's full trail out of the aggregator. Logging runs everywhere — including startup code, background jobs and panics recovered outside any request — so the lookup must degrade to a placeholder instead of panicking inside the logger.

## Task

Implement the exported function(s) in [requestid.go](requestid.go) so that:

1. `WithRequestID` stores `id` under the unexported key.
2. `RequestID` returns the stored string.
3. When the key is absent, holds a non-string, or holds the empty string, it returns `"unknown"`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  RequestID(WithRequestID(bg, "req-8f21"))
Output: "req-8f21"
```

**Example 2:**

```
Input:  RequestID(context.Background())
Output: "unknown"
```

**Example 3:**

```
Input:  RequestID(WithRequestID(bg, ""))
Output: "unknown"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`ctx.Value` lookup** | Returns `any`; the caller decides what a miss means. |
| 2 | **Comma-ok assertion** | Guards against both a missing key and a wrong-typed value. |
| 3 | **Safe defaults** | Observability code must never panic on missing metadata. |

## Hint

Assert with comma-ok, then treat `!ok` and `id == ""` the same way.

## Validate

```bash
make verify
```
