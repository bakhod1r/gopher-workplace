# Ordering That Ignores The Length Rule

**Level:** staff  
**Topic:** 03-generics

## Context

A router ranks candidate paths and is documented to prefer the shortest one, breaking ties by segment id. It keeps choosing a long path over a short one whenever the long path starts with a small id.

## Task

Fix the single planted bug in [stdcomparelenbug.go](stdcomparelenbug.go):

1. Find and fix the single bug so length dominates the ordering.
2. Paths of equal length must still compare segment by segment.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  ComparePaths([9], [1,2])
Output: -1
```

**Example 2:**

```
Input:  ComparePaths([1,2], [1,3])
Output: -1
```

**Example 3:**

```
Input:  ComparePaths([1,2], [1,2])
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Read the helper's ordering, not its name** | `slices.Compare` is lexicographic: it compares element-wise and only falls back to length when one slice is a prefix of the other. |
| 2 | **Composite orderings** | A two-level ordering is a length test first, then a delegated tie-break. |

## Hint

`slices.Compare([9], [1,2])` — work out what it returns and why.

## Validate

```bash
make verify
```
