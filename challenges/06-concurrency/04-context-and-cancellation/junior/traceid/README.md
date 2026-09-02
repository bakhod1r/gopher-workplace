# Propagating a Trace ID

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A gRPC interceptor reads the `traceparent` header off every inbound call and must make the trace ID reachable by the outbound clients several layers below, without changing the signature of every business function in between. That is exactly what request-scoped context values are for.

## Task

Implement the exported function(s) in [traceid.go](traceid.go) so that:

1. `WithTraceID` returns `context.WithValue(ctx, traceKey{}, id)`.
2. `TraceID` type-asserts `ctx.Value(traceKey{})` to `string` and returns the comma-ok result.
3. A missing trace yields `("", false)`; a trace explicitly set to the empty string yields `("", true)`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  TraceID(WithTraceID(context.Background(), "4bf92f"))
Output: "4bf92f", true
```

**Example 2:**

```
Input:  TraceID(context.Background())
Output: "", false
```

**Example 3:**

```
Input:  TraceID(WithTraceID(WithTraceID(bg, "a"), "b"))
Output: "b", true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.WithValue`** | Returns a new context carrying one key/value pair; the parent is unchanged. |
| 2 | **Unexported key types** | `type traceKey struct{}` makes collisions with other packages impossible. |
| 3 | **Comma-ok type assertion** | `v, ok := x.(string)` distinguishes "no trace" from "empty trace". |

## Hint

Use an unexported struct key. In `TraceID`, assert with the two-value form so a missing key gives `ok == false`.

## Validate

```bash
make verify
```
