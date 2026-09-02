# Concurrent Map

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Go's built-in map is not safe for concurrent use — the runtime deliberately
crashes on a detected concurrent write. Wrapping it in a struct with an
`RWMutex` gives many simultaneous readers or one exclusive writer.

## Task

Implement `Get` and `Set` on `*ConcurrentMap` in [concurmap.go](concurmap.go):

1. `Get` takes the **read** lock and returns the value and the comma-ok flag.
2. `Set` takes the **write** lock and stores the value.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Set("k", 1); Get("k")
Output: (1, true)
```

**Example 2:**

```
Input:  Get("missing")
Output: (0, false)
```

**Example 3:**

```
Input:  100 goroutines each Set then Get on the same key
Output: key exists; clean under `go test -race`
```

_Explanation:_ the last writer wins; which one is unspecified.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`RWMutex` read vs write** | `RLock` may be held by many goroutines; `Lock` excludes everyone. |
| 2 | **`defer` unlock** | Guarantees release on every return path, including panics. |
| 3 | **Comma-ok through a wrapper** | Return both results of the map lookup so callers keep the distinction. |

## Hint

Return the map lookup directly: `v, ok := m.data[key]; return v, ok`. Using
`Lock` in `Get` is correct but throws away all read concurrency.

## Validate

```bash
make verify
```
