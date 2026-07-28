# Mirror a Tree

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Mirroring swaps each node's children and recurses into both subtrees, mutating
the tree in place.

## Task

Implement `Mirror` in [mirrortree.go](mirrortree.go).

Do **not** change the function signature or the tests.

## Examples

```go
Mirror(1 with 2,3) // children become 3,2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Swap children** | `t.Left, t.Right = t.Right, t.Left`. |
| 2 | **Recurse both** | Mirror each subtree. |
| 3 | **nil base** | Nothing to do for nil. |

## Hint

`if t == nil { return }; t.Left, t.Right = t.Right, t.Left; Mirror(t.Left); Mirror(t.Right)`.

## Validate

```bash
make verify
```
