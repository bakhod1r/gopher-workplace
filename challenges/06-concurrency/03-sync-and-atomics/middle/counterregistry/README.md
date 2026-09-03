# Metrics Counter Registry

**Level:** middle
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A service registers a counter the first time a metric name is used, then bumps it on every event. Registration is rare; increments happen millions of times a second, so the hot path must not take a write lock.

## Task

Implement the stubbed functions in [counterregistry.go](counterregistry.go) so that:

1. `NewRegistry` initialises the map.
2. `Add` bumps a counter, registering the name on first use, and returns the new value.
3. An already-registered name must be served under the **read** lock only.
4. `Value` returns 0 for unknown names; `Snapshot` returns an independent copy.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  r.Add("http_requests", 1)
Output: 1
```

**Example 2:**

```
Input:  r.Add("http_requests", 2)
Output: 3
```

**Example 3:**

```
Input:  NewRegistry().Value("unknown")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Double-checked locking | Try under `RLock`; on a miss take `Lock` and **re-check** — another goroutine may have registered it meanwhile. |
| 2 | Pointer values in a map | Storing `*atomic.Int64` keeps the counter stable while the map entry is only read. |
| 3 | Lock scope vs atomics | The map needs the RWMutex; the counter itself needs none — `Add` is atomic on its own. |
| 4 | Copy under the lock | `Snapshot` must `Load` each counter while still holding `RLock`. |

## Hint

Factor out a `counter(name) *atomic.Int64` helper: `RLock` fast path, then `Lock` + re-check + create. `Add` is then a one-liner.

## Validate

```bash
make verify
```
