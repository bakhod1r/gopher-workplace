# Composite Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A folder holds files and other folders. The composite pattern makes the
container answer the same question as its parts: asking a folder for its size
gives one number, no matter how deep the tree goes.

## Task

Implement `Size` on `*Folder` in [compositepatt.go](compositepatt.go):

1. Sum every entry in `f.Files`.
2. Add `Size()` of every folder in `f.Sub`.
3. Return the total.

**Constraint (senior):** a 1000-deep tree of 100 files per level must be summed in one traversal, with no intermediate node list.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  &Folder{Files: []int{10, 20}}
Output: 30
```

**Example 2:**

```
Input:  &Folder{Sub: []*Folder{{Files: []int{30}}}}
Output: 30
```

**Example 3:**

```
Input:  &Folder{Files: []int{10,20}, Sub: []*Folder{{Files: []int{30}}, {Files: []int{40,50}}}}
Output: 150
```

_Explanation:_ 30 from own files plus 30 and 90 from the two subfolders.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recursive method** | A method may call itself on another receiver — the recursion rides on the tree, not on a helper function. |
| 2 | **Base case by construction** | A folder with no `Sub` ends the recursion; no explicit termination check is needed. |
| 3 | **`range` over slices** | Two loops, one accumulator. |

## Hint

One accumulator, two `range` loops: the second calls `sub.Size()`. An empty
`Sub` slice makes the second loop a no-op, which is the base case.

## Validate

```bash
make verify
```
