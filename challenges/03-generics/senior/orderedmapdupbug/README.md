# Duplicate Keys In The Order List

**Level:** senior  
**Topic:** 03-generics

## Context

A config renderer emits some fields twice — always the ones that were overridden after being set.

## Task

Fix the single planted bug in [orderedmapdupbug.go](orderedmapdupbug.go):

1. Find and fix the single bug so each key appears exactly once, at its first position.
2. Updating an existing key must keep its original position and its new value.
3. Only `Set` may change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Set(a,1); Set(b,2); Keys()
Output: [a b]
```

**Example 2:**

```
Input:  Set(a,1); Set(a,2); Keys()
Output: [a]
```

**Example 3:**

```
Input:  Set(a,1); Set(a,2); Get(a)
Output: 2, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two structures, one invariant** | `len(keys)` must equal `len(items)` after every operation. |
| 2 | **Insert versus update** | Only a genuinely new key extends the order list. |
| 3 | **Cheap check** | Comparing the two lengths in a test catches this class of bug immediately. |

## Hint

What does the second `Set` on the same key do to `keys`?

## Validate

```bash
make verify
```
