# Session Store

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A web server keeps logged-in sessions in memory: a session token maps to a user ID. Login handlers write and every authenticated request reads, all concurrently. A bare Go map aborts the process with `fatal error: concurrent map writes`.

## Task

Implement the stubbed functions in [sessionstore.go](sessionstore.go) so that:

1. `NewSessionStore` returns a ready store with an initialised map.
2. `Save` stores the user ID for a session token under the lock.
3. `Lookup` returns the user ID and whether the session exists.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  s := NewSessionStore(); s.Save("tok1", 7); s.Lookup("tok1")
Output: 7, true
```

**Example 2:**

```
Input:  s := NewSessionStore(); s.Lookup("unknown")
Output: 0, false
```

**Example 3:**

```
Input:  s.Save("tok1", 7); s.Save("tok1", 9); s.Lookup("tok1")
Output: 9, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Mutex** | One lock guards the whole session map. |
| 2 | **Map comma-ok** | `v, ok := m[k]` reports presence - an unknown token must not look like user 0. |
| 3 | **Constructor** | Maps must be made before use; writing to a nil map panics. |

## Hint

One `sync.Mutex` field guards the map. Lock in both `Save` and `Lookup`.

## Validate

```bash
make verify
go test -race ./...
```
