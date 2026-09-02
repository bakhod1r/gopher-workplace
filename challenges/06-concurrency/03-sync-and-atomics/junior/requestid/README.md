# Request ID Generator

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

Every inbound request gets a unique, increasing ID for the trace logs. Handlers on different cores ask for IDs at the same instant, and two requests sharing an ID would make the traces impossible to read.

## Task

Implement the stubbed functions in [requestid.go](requestid.go) so that:

1. `Next` returns the next ID, starting at 1 and never repeating.
2. `Issued` reports how many IDs have been handed out.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var g IDGen; g.Next()
Output: 1
```

**Example 2:**

```
Input:  var g IDGen; g.Next(); g.Next()
Output: 1, then 2
```

**Example 3:**

```
Input:  var g IDGen; g.Next(); g.Issued()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Int64.Add** | `Add` returns the *new* value — that returned value is your unique ID. |
| 2 | **Uniqueness** | Read-then-write hands two callers the same number; one atomic add cannot. |
| 3 | **Load vs Add** | `Issued` observes without incrementing. |

## Hint

`return g.n.Add(1)` — a single call both increments and gives you the value nobody else got.

## Validate

```bash
make verify
go test -race ./...
```
