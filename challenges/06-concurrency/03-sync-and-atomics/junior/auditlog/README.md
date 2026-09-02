# Audit Log Buffer

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A compliance service buffers audit entries in memory before flushing them to storage. Handlers append from every goroutine, and the flusher reads the whole buffer — which must never be the same backing array the handlers are still appending to.

## Task

Implement the stubbed functions in [auditlog.go](auditlog.go) so that:

1. `Append` adds an entry under the lock.
2. `Entries` returns a **copy** of the buffered entries, in order.
3. `Len` returns how many entries are buffered.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var l AuditLog; l.Append("login"); l.Len()
Output: 1
```

**Example 2:**

```
Input:  l.Append("login"); l.Append("logout"); l.Entries()
Output: ["login", "logout"]
```

**Example 3:**

```
Input:  var l AuditLog; l.Entries()
Output: [] (empty)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slice aliasing** | A returned slice shares its backing array — copy it inside the lock. |
| 2 | **append under lock** | `append` may reallocate; two concurrent appends corrupt the slice header. |
| 3 | **sync.Mutex** | One lock guards the buffer for both append and read. |

## Hint

In `Entries`, `out := make([]string, len(l.entries))`, then `copy(out, l.entries)`, and return `out`.

## Validate

```bash
make verify
go test -race ./...
```
