# Server Root Context

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The checkout API's `main` builds one root context at process start and threads it into the HTTP server, the Postgres pool and the OpenTelemetry exporter. Every per-request and per-query context in the service is derived from it, so it must be the empty root: never cancelled, no deadline, no values.

## Task

Implement the exported function(s) in [servercontext.go](servercontext.go) so that:

1. It returns `context.Background()`.
2. The returned context has a `nil` `Done()` channel and a `nil` `Err()`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ServerContext() != nil
Output: true
```

**Example 2:**

```
Input:  ServerContext().Err()
Output: nil
```

**Example 3:**

```
Input:  ServerContext().Done()
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.Background()`** | The empty root context: never cancelled, no deadline, no values. |
| 2 | **`Done()` may be nil** | A context that can never be cancelled returns a nil channel; receiving from it blocks forever. |
| 3 | **`Err()` contract** | `Err()` is nil while the context is alive. |

## Hint

`context.Background()` is a package-level singleton — just return it.

## Validate

```bash
make verify
```
