# First Occurrence Not Kept

**Level:** senior  
**Topic:** 03-generics

## Context

An import feed keeps the wrong version of each record: re-sent rows with stale timestamps overwrite the good ones.

## Task

Fix the single planted bug in [uniquekeybug.go](uniquekeybug.go):

1. Find and fix the single bug so the **first** element for each key survives.
2. Input order must be preserved.
3. The element type must stay unconstrained.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  UniqueBy([{1 a} {1 b} {2 c}], idOf)
Output: [{1 a} {2 c}]
```

**Example 2:**

```
Input:  UniqueBy([], idOf)
Output: []
```

**Example 3:**

```
Input:  order
Output: first-seen
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **First versus last wins** | Assigning unconditionally into a map keeps the last value per key. |
| 2 | **Order is not the only contract** | The buggy version preserves order and still returns the wrong elements. |
| 3 | **Simplest correct shape** | A `seen` set plus a single append avoids the question entirely. |

## Hint

The output order is right. Which *element* does each key end up holding?

## Validate

```bash
make verify
```
