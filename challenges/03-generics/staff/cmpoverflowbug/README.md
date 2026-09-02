# The Comparator That Wraps Around

**Level:** staff  
**Topic:** 03-generics

## Context

A ledger sorts entries by a signed balance in the smallest units. Ordinary balances sort fine; accounts holding sentinel values near the extremes of int sort to the wrong end, and reconciliation reports the wrong worst offender.

## Task

Fix the single planted bug in [cmpoverflowbug.go](cmpoverflowbug.go):

1. Find and fix the single bug so extreme keys compare correctly.
2. The sort must stay in place and keep the same signature.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  keys {MaxInt, MinInt}
Output: MinInt first
```

**Example 2:**

```
Input:  keys {0, MaxInt, MinInt, 5, -5}
Output: MinInt, -5, 0, 5, MaxInt
```

**Example 3:**

```
Input:  keys {3, 1, 2}
Output: 1, 2, 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Signed overflow wraps silently** | `a - b` on int wraps modulo 2^64; the sign of the result is then meaningless. |
| 2 | **Compare, do not subtract** | `cmp.Compare` branches on the operands and never overflows. |
| 3 | **Broken comparators corrupt sorts** | A comparator that is not a valid ordering leaves the slice arbitrarily scrambled, not merely imperfect. |

## Hint

The comparator calls `cmp.Compare`. Look at what it is comparing.

## Validate

```bash
make verify
```
