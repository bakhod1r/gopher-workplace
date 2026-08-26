# Tree Walk

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A binary search tree needs to produce a sorted list of its elements via
in-order traversal.

## Task

Implement `Walk` on `*Tree` in [treewalk.go](treewalk.go):

1. Return in-order traversal: left subtree, root, right subtree.
2. Return `[]int{}` (not nil) for a nil tree.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Tree{2, Tree{1}, Tree{3}}.Walk()
Output: [1, 2, 3]
```

**Example 2:**

```
Input:  (*Tree)(nil).Walk()
Output: []
```

**Example 3:**

```
Input:  Tree{3, Tree{2, Tree{1}, nil}, nil}.Walk()
Output: [1, 2, 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil receiver** | A nil `*Tree` should return empty, not panic. |
| 2 | **Recursion** | Call `Walk` on left and right subtrees. |
| 3 | **Pointer receiver** | `*Tree` to handle nil. |

## Hint

If `t == nil`, return `[]int{}`. Otherwise: `left.Walk() + [t.Val] + right.Walk()`.

## Validate

```bash
make verify
```
