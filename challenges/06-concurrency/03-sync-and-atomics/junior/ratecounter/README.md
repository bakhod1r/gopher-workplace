# Rate Counter

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An API gateway counts inbound requests so it can report requests-per-second. Every accepted request is recorded from the goroutine that served it, so the counter is written from dozens of goroutines at once. A plain `hits++` loses updates and trips the race detector.

## Task

Implement the stubbed functions in [ratecounter.go](ratecounter.go) so that:

1. `Record` counts one request, safely under concurrent calls.
2. `Hits` returns the number of requests recorded so far.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var c RateCounter; c.Record(); c.Hits()
Output: 1
```

**Example 2:**

```
Input:  var c RateCounter; c.Record(); c.Record(); c.Hits()
Output: 2
```

**Example 3:**

```
Input:  var c RateCounter; c.Hits()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Mutex** | `Lock`/`Unlock` make a critical section exclusive. |
| 2 | **Pointer receiver** | A mutex must not be copied - methods take `*RateCounter`. |
| 3 | **Zero value ready** | `sync.Mutex`'s zero value is an unlocked mutex; no constructor needed. |

## Hint

Wrap every access to `c.hits` - reads included - in `c.mu.Lock()` / `c.mu.Unlock()`.

## Validate

```bash
make verify
go test -race ./...
```
