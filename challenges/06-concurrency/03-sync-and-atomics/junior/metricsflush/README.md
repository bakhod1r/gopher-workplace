# Metrics Flush

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A telemetry agent accumulates a byte counter and ships it to the collector every ten seconds. The flush must read the total *and* zero it in one step — anything counted between the read and the reset would be lost.

## Task

Implement the stubbed functions in [metricsflush.go](metricsflush.go) so that:

1. `Record` adds a value to the pending counter.
2. `Drain` returns the pending total and resets it to zero, atomically.
3. `Pending` returns the current total without clearing it.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var m Meter; m.Record(3); m.Drain()
Output: 3, pending now 0
```

**Example 2:**

```
Input:  var m Meter; m.Record(3); m.Record(4); m.Pending()
Output: 7
```

**Example 3:**

```
Input:  var m Meter; m.Drain()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Int64.Swap** | `Swap(0)` stores 0 and returns the previous value in one indivisible step. |
| 2 | **Lost updates** | `Load` then `Store(0)` drops everything recorded in between. |
| 3 | **Read-and-reset** | The classic counter-flush idiom in metrics agents. |

## Hint

`return m.n.Swap(0)` — one call reads the old total and clears it.

## Validate

```bash
make verify
go test -race ./...
```
