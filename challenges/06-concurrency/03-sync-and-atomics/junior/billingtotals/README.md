# Billing Run Totals

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A nightly billing run sums invoice amounts across many worker goroutines and also counts how many invoices were included. The sum and the count must always agree — a report that says 3 invoices totalling the amount of 4 is worse than no report at all.

## Task

Implement the stubbed functions in [billingtotals.go](billingtotals.go) so that:

1. `Add` records one invoice amount, updating the sum and the count together.
2. `Total` and `Count` return the accumulated values.
3. `Average` returns the mean amount, truncated, and 0 when no invoices were added.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var t Totals; t.Add(300); t.Total()
Output: 300
```

**Example 2:**

```
Input:  var t Totals; t.Add(300); t.Add(100); t.Average()
Output: 200
```

**Example 3:**

```
Input:  var t Totals; t.Average()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Multi-field invariant** | Two fields that must agree are updated inside one lock hold. |
| 2 | **sync.Mutex** | The lock protects a relationship, not just individual variables. |
| 3 | **Integer division** | `sum / count` truncates; guard `count == 0` first. |

## Hint

`Add` locks once and updates both `sum` and `count`; `Average` locks once and reads both.

## Validate

```bash
make verify
go test -race ./...
```
