# Feature Flag Set

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

Feature flags are read on nearly every request and rewritten only when an operator pushes a new configuration. A plain `Mutex` would serialise those thousands of reads behind each other; an `RWMutex` lets them proceed in parallel and only blocks them during the rare write.

## Task

Implement the stubbed functions in [flagset.go](flagset.go) so that:

1. `Set` takes the write lock and records a flag's state.
2. `Enabled` takes the read lock and reports whether a flag is on (unknown flags are off).
3. `Len` reports how many flags are configured, under the read lock.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  f := NewFlagSet(); f.Set("new_ui", true); f.Enabled("new_ui")
Output: true
```

**Example 2:**

```
Input:  f := NewFlagSet(); f.Enabled("unknown")
Output: false
```

**Example 3:**

```
Input:  f.Set("a", true); f.Set("b", false); f.Len()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.RWMutex** | `RLock` allows many concurrent readers; `Lock` excludes everyone. |
| 2 | **Read-mostly workloads** | Use `RWMutex` when reads vastly outnumber writes. |
| 3 | **Lock pairing** | `RLock` pairs with `RUnlock`, `Lock` with `Unlock` - never mix them. |

## Hint

The operator path calls `mu.Lock()`/`mu.Unlock()`; the request path calls `mu.RLock()`/`mu.RUnlock()`.

## Validate

```bash
make verify
go test -race ./...
```
