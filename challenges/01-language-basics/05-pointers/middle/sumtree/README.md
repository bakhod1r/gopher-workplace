# Sum a Tree

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Summing a tree adds the node value to the sums of both subtrees.

## Task

Implement `SumTree` in [sumtree.go](sumtree.go).

Do **not** change the function signature or the tests.

## Examples

```go
SumTree(1 with 2,3 children) // => 6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Recursive combine** | val + left + right. |
| 2 | **nil base** | 0 for nil. |
| 3 | **Post-order shape** | Children before combining. |

## Hint

`if t == nil { return 0 }; return t.Val + SumTree(t.Left) + SumTree(t.Right)`.

## Validate

```bash
make verify
```
