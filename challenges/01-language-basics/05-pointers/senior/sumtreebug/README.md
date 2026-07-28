# Tree Sum Skips Subtrees

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Summing a tree must add the two subtree sums to the node value. Returning only
`t.Val` ignores every descendant.

## Task

Fix [sumtreebug.go](sumtreebug.go) to recurse into both children.

Do **not** change the function signature or the tests.

## Examples

```go
SumTree(tree of 1,2,3,4) // => 10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Recursive combine** | val + left sum + right sum. |
| 2 | **Both children** | Recurse Left AND Right. |
| 3 | **Post-order** | Children before combining. |

## Hint

Recurse: `return t.Val + SumTree(t.Left) + SumTree(t.Right)`.

## Validate

```bash
make verify
```
