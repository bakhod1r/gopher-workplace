# Insert Overwrites the Root

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Reassigning `root = &Tree{Val: v}` throws away the existing tree each call. A
real insert descends to a nil slot and creates a node there, leaving the rest
intact.

## Task

Fix [insertoverwrite.go](insertoverwrite.go) to insert without destroying the tree.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Insert 5, then 3, then 8
Output: 3-node BST
```

**Example 2:**

```
Input:  Insert(nil, 5)
Output: root 5
```

**Example 3:**

```
Input:  Insert(root(5), 3)
Output: 3 as left child
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Preserve the tree** | Only create a node at a nil slot. |
| 2 | **Recursive descent** | Recurse into the correct child. |
| 3 | **Return the root** | Reattach child subtrees. |

## Hint

Descend and insert: `if root == nil { return &Tree{Val: v} }; if v < root.Val { root.Left = Insert(root.Left, v) } else { root.Right = Insert(root.Right, v) }; return root`.

## Validate

```bash
make verify
```
