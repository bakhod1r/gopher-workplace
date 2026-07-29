# BST Search

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

BST search prunes half the tree at each step: go left when the target is
smaller, right when larger.

## Task

Implement `Contains` in [bstcontains.go](bstcontains.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Contains(bst, 8)
Output: true
```

**Example 2:**

```
Input:  Contains(bst, 4)
Output: false
```

**Example 3:**

```
Input:  Contains(nil, 1)
Output: false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Ordered descent** | Compare v with node value. |
| 2 | **Prune** | Only one subtree can hold v. |
| 3 | **nil base** | Not found at nil. |

## Hint

`if root == nil { return false }; if v == root.Val { return true }; if v < root.Val { return Contains(root.Left, v) }; return Contains(root.Right, v)`.

## Validate

```bash
make verify
```
