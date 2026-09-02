# Lazy Database Pool

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A service opens its database connection pool lazily, on the first query rather than at start-up. Dozens of request goroutines may ask for the pool simultaneously, but opening it twice would double the connection count and blow the server's `max_connections`.

## Task

Implement the stubbed functions in [dbpoolinit.go](dbpoolinit.go) so that:

1. `Pool` opens the pool on the first call and caches it.
2. Every later call returns the same `*Pool` without opening another.
3. `Opens` reports how many times the open function actually ran.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  p := NewProvider(func() *Pool { return &Pool{DSN: "db"} }); p.Pool().DSN
Output: "db"
```

**Example 2:**

```
Input:  p.Pool(); p.Pool(); p.Opens()
Output: 1
```

**Example 3:**

```
Input:  p := NewProvider(...); p.Opens()  // before any Pool call
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once** | `once.Do(f)` runs `f` exactly once, ever, and blocks other callers until it finishes. |
| 2 | **Closures** | The open function is a `func() *Pool` captured by the provider. |
| 3 | **Lazy initialisation** | Expensive set-up is deferred until first use, then memoised. |

## Hint

Keep a `sync.Once`, the cached `*Pool`, and a counter. Do the open *and* the counter increment inside `once.Do`.

## Validate

```bash
make verify
go test -race ./...
```
