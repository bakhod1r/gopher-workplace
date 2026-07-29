# Insert Result Not Reattached

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

When a subtree is nil, `Insert` returns a NEW node that the parent must reattach
via `root.Left = Insert(root.Left, v)`. Discarding the return drops the new node.

## Task

Fix [insertnoreattach.go](insertnoreattach.go) so left-subtree inserts persist.

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
Output: 3 attached as left child
```

**Example 3:**

```
Input:  Insert(root(5), 8)
Output: 8 attached as right child
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Reattach the subtree** | `root.Left = Insert(root.Left, v)`. |
| 2 | **New nodes returned** | Empty slots return a fresh node. |
| 3 | **Symmetry** | Both children need reassignment. |

## Hint

Assign the result back: `root.Left = Insert(root.Left, v)`.

## Validate

```bash
make verify
```
