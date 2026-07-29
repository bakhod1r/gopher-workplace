# Tree Height

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Tree height is 1 plus the taller of the two subtrees; nil subtrees contribute 0.

## Task

Implement `Height` in [treeheight.go](treeheight.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Height(nil)
Output: 0
```

**Example 2:**

```
Input:  Height(single node)
Output: 1
```

**Example 3:**

```
Input:  Height(balanced 3 levels)
Output: 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Recursive over children** | Recurse Left and Right. |
| 2 | **Combine** | `1 + max(left, right)`. |
| 3 | **nil base** | nil is height 0. |

## Hint

`if t == nil { return 0 }; l, r := Height(t.Left), Height(t.Right); if l > r { return 1+l }; return 1+r`.

## Validate

```bash
make verify
```
