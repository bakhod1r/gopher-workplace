# Cache Hit Ratio

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A CDN edge node reports its cache hit ratio. Every request records a hit or a miss from its own goroutine, and the metrics endpoint reads both counters to compute the ratio.

## Task

Implement the stubbed functions in [cachestats.go](cachestats.go) so that:

1. `Hit` and `Miss` each count one lookup atomically.
2. `Hits` and `Misses` return the two counts.
3. `Ratio` returns hits divided by total lookups, and 0 when there have been none.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var s Stats; s.Hit(); s.Hits()
Output: 1
```

**Example 2:**

```
Input:  var s Stats; s.Hit(); s.Miss(); s.Ratio()
Output: 0.5
```

**Example 3:**

```
Input:  var s Stats; s.Ratio()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Int64** | Two independent counters, each updated with an indivisible `Add`. |
| 2 | **Division guard** | Total zero means no lookups — return 0 rather than NaN. |
| 3 | **Snapshot skew** | Two separate `Load`s are each atomic, but the pair is not a single instant. |

## Hint

`Ratio` loads both counters once into locals, guards `total == 0`, then divides the two `float64` conversions.

## Validate

```bash
make verify
go test -race ./...
```
