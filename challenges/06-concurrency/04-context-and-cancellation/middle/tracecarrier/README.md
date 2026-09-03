# Trace ID Carrier

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Every request entering the API gateway gets a trace ID that must reach the database driver, the cache client, and the audit logger without appearing in any function signature. The value travels on the request context, keyed by an unexported type so no other package can collide with it.

## Task

Implement the stubbed functions in [tracecarrier.go](tracecarrier.go) so that:

1. `WithTrace` attaches the ID under the unexported `traceKey{}`; an empty ID returns the parent unchanged.
2. `TraceID` returns the ID, or `""` when the context carries none.
3. A plain `string` key from another package must never be mistaken for the trace key.
4. `Chain` reads the ID before tagging, then twice after.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  TraceID(WithTrace(ctx, "abc"))
Output: "abc"
```

**Example 2:**

```
Input:  TraceID(context.Background())
Output: ""
```

**Example 3:**

```
Input:  Chain(ctx, "abc")
Output: ["" "abc" "abc"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Unexported key type | `type traceKey struct{}` cannot be constructed by another package, so keys never collide. |
| 2 | Comma-ok type assertion | `ctx.Value` returns `any` and `nil` when absent — assert with the two-value form and the zero value falls out. |
| 3 | Context values are immutable | `WithValue` returns a *new* context; the parent never changes. |
| 4 | Request-scoped data only | Trace IDs and auth principals belong here; optional parameters and dependencies do not. |

## Hint

`context.WithValue(ctx, traceKey{}, id)` to write, `ctx.Value(traceKey{}).(string)` with comma-ok to read.

## Validate

```bash
make verify
```
