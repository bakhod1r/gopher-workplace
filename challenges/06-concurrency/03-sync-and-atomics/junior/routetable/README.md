# Gateway Routing Table

**Level:** junior
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An API gateway looks up a backend for every inbound request, while a config watcher rewrites individual routes during a deploy. Reads outnumber writes by orders of magnitude, and an admin endpoint dumps the whole table without blocking traffic.

## Task

Implement the stubbed functions in [routetable.go](routetable.go) so that:

1. `NewTable` returns an initialised, empty table.
2. `Set` writes a route under the write lock; `Lookup` reads under the read lock.
3. `Snapshot` returns a **copy** the caller can read and even mutate without touching the table.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  t.Set("/api", "backend-1"); t.Lookup("/api")
Output: "backend-1", true
```

**Example 2:**

```
Input:  NewTable().Lookup("/missing")
Output: "", false
```

**Example 3:**

```
Input:  t.Snapshot()
Output: map[/api:backend-1]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.RWMutex` | Many readers may hold `RLock` at once; a writer's `Lock` excludes everyone. |
| 2 | Read-mostly data | When reads dominate, `RWMutex` beats a plain `Mutex` — but only if writes are genuinely rare. |
| 3 | Copy under the lock | Returning the internal map hands the caller an unsynchronised alias. Copy it while the lock is held. |

## Hint

`Lookup` and `Snapshot` take `RLock`; `Set` takes `Lock`. Build the snapshot's map inside the critical section, not after it.

## Validate

```bash
make verify
```
