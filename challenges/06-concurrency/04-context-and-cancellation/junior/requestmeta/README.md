# Tenant and Trace Together

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The multi-tenant API attaches two pieces of metadata to every inbound request: the tenant ID, which the data layer uses to scope every query, and the trace ID, which the logger and the outbound clients propagate. Both ride along on the context so no intermediate signature has to mention them.

## Task

Implement the exported function(s) in [requestmeta.go](requestmeta.go) so that:

1. `WithMeta` stores the tenant under `tenantKey{}` and the trace under `traceKey{}`, wrapping twice.
2. `Meta` looks up both keys and returns them.
3. Missing values come back as `""`, and the innermost `WithMeta` wins.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Meta(WithMeta(bg, "acme", "4bf9"))
Output: "acme", "4bf9"
```

**Example 2:**

```
Input:  Meta(context.Background())
Output: "", ""
```

**Example 3:**

```
Input:  Meta(WithMeta(WithMeta(bg, "a", "1"), "b", "2"))
Output: "b", "2"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **One key per value** | `WithValue` stores a single pair; two values means two wrappers. |
| 2 | **Distinct key types** | `tenantKey{}` and `traceKey{}` are different types, so they never collide. |
| 3 | **Assertion with a discarded ok** | `v, _ := x.(string)` yields `""` on a miss, which is the wanted default here. |

## Hint

Chain the two `WithValue` calls: wrap for the tenant, then wrap that result for the trace.

## Validate

```bash
make verify
```
