# The Update That Half Survives

**Level:** staff  
**Topic:** 03-generics

## Context

An analytics fan-out updates several tallies per event. The per-value breakdown looks right, but every tally's `Total` stays at zero, so every percentage is computed as a division by zero.

## Task

Fix the single planted bug in [tallycopybug.go](tallycopybug.go):

1. Find and fix the single bug so both the map and the counter advance.
2. The tallies must stay consistent after any number of bumps.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  BumpAll(ts, "x"); ts[0].Counts["x"]
Output: 1
```

**Example 2:**

```
Input:  BumpAll(ts, "x"); ts[0].Total
Output: 1
```

**Example 3:**

```
Input:  two bumps; Consistent(ts[0])
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Range copies the element** | `for _, t := range ts` hands you a copy of the struct, not the element. |
| 2 | **Reference fields survive the copy** | The copy's map header points at the same table, so map writes stick while plain-field writes do not. |
| 3 | **Index to mutate** | `for i := range ts` plus `ts[i]` writes through to the slice. |

## Hint

One of the two updates lands and the other does not. Ask why they differ.

## Validate

```bash
make verify
```
