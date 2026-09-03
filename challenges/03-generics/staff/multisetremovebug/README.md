# Remove That Empties The Bin

**Level:** staff  
**Topic:** 03-generics

## Context

An inventory multiset tracks how many of each SKU are on hand. Shipping one unit of a SKU that had twelve in stock drops the on-hand count to zero, and the totals stop matching the ledger.

## Task

Fix the single planted bug in [multisetremovebug.go](multisetremovebug.go):

1. Find and fix the single bug so removing one occurrence decrements the tally.
2. The key must disappear only when its last occurrence goes.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add x three times; Remove(x); Count(x)
Output: 2
```

**Example 2:**

```
Input:  Add x three times; Remove(x); Distinct()
Output: 1
```

**Example 3:**

```
Input:  Add x once; Remove(x); Remove(x)
Output: true then false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Multiplicity is the point** | A multiset is a map of counts; deleting the key throws the multiplicity away. |
| 2 | **Two totals, one operation** | The per-key count and the running total must move together. |
| 3 | **Delete at zero, not before** | The key is retired only when its count reaches zero. |

## Hint

What does `delete` do to a count of twelve?

## Validate

```bash
make verify
```
