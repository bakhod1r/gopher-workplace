# Connection Registry

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A service mesh sidecar keeps a registry of live upstream connections: an instance ID mapped to its address. Routing goroutines read it constantly; the discovery loop rewrites it every few seconds. Reads must not block each other, and a listing must never hand out the registry's own storage.

## Task

Implement the stubbed functions in [connregistry.go](connregistry.go) so that:

1. `Register` records an instance's address under the write lock.
2. `Lookup` returns an instance's address and whether it is registered.
3. `IDs` returns a **copy** of the registered instance IDs, in any order.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  r := NewRegistry(); r.Register("a", "10.0.0.1"); r.Lookup("a")
Output: "10.0.0.1", true
```

**Example 2:**

```
Input:  r := NewRegistry(); r.Lookup("ghost")
Output: "", false
```

**Example 3:**

```
Input:  r.Register("b", "x"); r.Register("a", "y"); len(r.IDs())
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.RWMutex** | Routing reads take `RLock`; discovery writes take `Lock`. |
| 2 | **Defensive copy** | Build a new slice inside the lock; never return internal state. |
| 3 | **Map iteration order** | Ranging a map is deliberately unordered; callers sort if they need order. |

## Hint

In `IDs`, take `RLock`, `make` a slice with `len(r.conns)` capacity, append every key, and return that fresh slice.

## Validate

```bash
make verify
go test -race ./...
```
