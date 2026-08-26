# Async Fetch

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A service needs to run operations asynchronously and signal when they complete.
A common pattern is a method returning a `<-chan struct{}`.

## Task

Implement `FetchAsync` on `*Fetcher` in [asyncfetch.go](asyncfetch.go):

1. Create a `chan struct{}`.
2. Launch a goroutine that calls `f.Fetch(id)`.
3. Close the channel when `Fetch` finishes.
4. Return the channel immediately.

Do **not** change the function signatures or the tests.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutines in methods** | Methods can launch goroutines just like functions. |
| 2 | **Completion channels** | `close(done)` signals all receivers. |

## Validate

```bash
make verify
```
