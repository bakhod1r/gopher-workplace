# Per-Collection Lazy Index

**Level:** middle  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

Each searchable collection needs a compiled index, and compiling one is expensive enough that doing it twice is a visible latency spike. Requests arrive in bursts, so thirty goroutines can ask for the same cold collection in the same millisecond. `sync.Once` gives exactly-once semantics — but a single `Once` covers a single thing, and here there is one thing *per collection*.

## Task

Implement the exported function(s) in [indexbuilder.go](indexbuilder.go) so that:

1. `NewRegistry` initialises both maps and stores the `build` func.
2. `Index` finds or creates the `*sync.Once` for the collection while holding the mutex, then releases the mutex.
3. `Index` calls `once.Do`, whose body runs `build` and stores the result under the mutex.
4. `Index` returns the compiled index; `build` runs exactly once per collection however many goroutines call in.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  r := NewRegistry(f); r.Index("orders")
Output: "orders-idx"
```

**Example 2:**

```
Input:  r.Index("orders"); r.Index("orders")
Output: build ran once
```

**Example 3:**

```
Input:  r.Index("orders"); r.Index("users"); r.Built()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once** | `Do` runs its func once and blocks every other caller until that run finishes. |
| 2 | **Once per key** | A map of `*sync.Once` scales exactly-once from one resource to many. |
| 3 | **Never hold a lock across slow work** | `build` runs inside `Do` but outside the registry mutex. |

## Hint

Two short critical sections, with `once.Do` between them: one to get-or-create the `*sync.Once`, one inside `Do` to store the result.

## Validate

```bash
make verify
go test -race ./...
```
