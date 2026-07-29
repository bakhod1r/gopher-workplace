# Count Leaf Nodes

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A leaf has no children. Count leaves by recursing and summing the leaf counts
of both subtrees.

## Task

Implement `CountLeaves` in [countleaves.go](countleaves.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountLeaves(nil)
Output: 0
```

**Example 2:**

```
Input:  CountLeaves(single node)
Output: 1
```

**Example 3:**

```
Input:  CountLeaves(root with 2 children)
Output: 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Leaf condition** | Both children nil. |
| 2 | **Recursive sum** | left + right leaves. |
| 3 | **nil base** | nil contributes 0. |

## Hint

`if t == nil { return 0 }; if t.Left == nil && t.Right == nil { return 1 }; return CountLeaves(t.Left) + CountLeaves(t.Right)`.

## Validate

```bash
make verify
```
