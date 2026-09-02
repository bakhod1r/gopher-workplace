# Per-Route Hit Counter

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An API gateway breaks its traffic down per route: `/users`, `/orders`, and so on. Every serving goroutine bumps the counter for the route it handled, so the whole map is written concurrently.

## Task

Implement the stubbed functions in [endpointhits.go](endpointhits.go) so that:

1. `NewHitCounter` returns a ready counter.
2. `Record` increments the count for a route, safely.
3. `Hits` returns the hits for a route (0 if the route was never served).

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  h := NewHitCounter(); h.Record("/users"); h.Hits("/users")
Output: 1
```

**Example 2:**

```
Input:  h.Record("/users"); h.Record("/users"); h.Hits("/users")
Output: 2
```

**Example 3:**

```
Input:  h := NewHitCounter(); h.Hits("/orders")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Mutex** | One lock protects the whole per-route map. |
| 2 | **Map increment** | `m[k]++` works on a missing key: it starts from the zero value. |
| 3 | **Critical section** | Read-modify-write on a map entry must be inside one lock hold. |

## Hint

`h.hits[route]++` is a read *and* a write; do it entirely inside a single `Lock`/`Unlock` pair.

## Validate

```bash
make verify
go test -race ./...
```
