# The Loop Variable Is A Copy

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A stats pass increments a counter in every element of a slice. It runs without error and the counters never move.

## Task

Fix the single planted bug in [rangecopy.go](rangecopy.go):

1. Increment `N` in every element of `items`, in place.
2. The writes must be visible through the caller's slice.
3. Fix the single bug; allocate nothing.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  items := []Counter{{N:1}}; Bump(items)
Output: items[0].N is 2
```

**Example 2:**

```
Input:  Bump(items[1:2])
Output: only the middle element changes
```

**Example 3:**

```
Input:  Bump(nil)
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **range copies the element** | The value variable is a fresh copy each iteration. |
| 2 | **Index to write** | `items[i]` addresses the slice's own storage. |
| 3 | **The copy can be expensive** | A 72-byte element is copied per iteration for nothing. |

## Hint

Which variable is the slice's element, and which is a copy of it?

## Validate

```bash
make verify
```
