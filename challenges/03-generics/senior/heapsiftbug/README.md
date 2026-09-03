# Heap That Loses Its Order

**Level:** senior  
**Topic:** 03-generics

## Context

A scheduler occasionally runs a later deadline first. It only happens when both children of a node are smaller than their parent.

## Task

Fix the single planted bug in [heapsiftbug.go](heapsiftbug.go):

1. Find and fix the single bug in `Pop` so the heap invariant survives every removal.
2. `Push` is correct — do not change it.
3. Popping every element must yield ascending order.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  push 5,3,8,1,9,2 then drain
Output: 1 2 3 5 8 9
```

**Example 2:**

```
Input:  Pop on empty
Output: zero, false
```

**Example 3:**

```
Input:  push 2,1 then Pop
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structural invariants** | A data structure is only correct if every operation restores what it promises. |
| 2 | **Sift-down compares against the current best** | Both children must be compared with the running smallest, not with the parent. |
| 3 | **Rare inputs expose it** | The bug only shows when the right child is smaller than the parent but larger than the left child. |

## Hint

Look closely at what the right-child comparison is measured against.

## Validate

```bash
make verify
```
