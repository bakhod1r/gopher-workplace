# BST Insert

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

BST insertion descends left for smaller values and right for larger/equal,
creating a node when it reaches a nil slot. Returning the root handles the
empty-tree case.

## Task

Implement `Insert` in [bstinsert.go](bstinsert.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Insert(nil, 5)
Output: tree rooted at 5
```

**Example 2:**

```
Input:  Insert(root(5), 3)
Output: 3 becomes left child
```

**Example 3:**

```
Input:  Insert(root(5), 8)
Output: 8 becomes right child
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **BST property** | left < node <= right. |
| 2 | **Recursive descent** | Recurse into the correct child. |
| 3 | **Return the subtree** | Reassign `t.Left`/`t.Right`. |

## Hint

`if root == nil { return &Tree{Val: v} }; if v < root.Val { root.Left = Insert(root.Left, v) } else { root.Right = Insert(root.Right, v) }; return root`.

## Validate

```bash
make verify
```
