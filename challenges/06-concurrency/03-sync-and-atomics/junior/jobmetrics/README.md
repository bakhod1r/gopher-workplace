# Job Metrics

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A background worker pool reports how many jobs it has processed. A mutex would work, but a single 64-bit counter hammered by every worker is exactly what `atomic.Int64` is for: no lock, no contention, no lost updates.

## Task

Implement the stubbed functions in [jobmetrics.go](jobmetrics.go) so that:

1. `Add` adds a delta to the processed count atomically.
2. `Processed` returns the current count.
3. `Reset` zeroes the counter after a metrics flush.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var m JobMetrics; m.Add(3); m.Processed()
Output: 3
```

**Example 2:**

```
Input:  var m JobMetrics; m.Add(2); m.Add(-5); m.Processed()
Output: -3
```

**Example 3:**

```
Input:  var m JobMetrics; m.Add(7); m.Reset(); m.Processed()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Int64** | `Add`, `Load` and `Store` are indivisible - no lock needed. |
| 2 | **Zero value ready** | `atomic.Int64`'s zero value is a zero counter. |
| 3 | **Pointer receiver** | Atomics must not be copied once used. |

## Hint

Hold an `atomic.Int64` field and forward to its `Add`, `Load` and `Store` methods.

## Validate

```bash
make verify
go test -race ./...
```
