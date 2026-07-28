# Count Nodes by Predicate

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Combining pointer traversal with a predicate function counts matching nodes in
one pass.

## Task

Implement `CountIf` in [countif.go](countif.go).

Do **not** change the function signature or the tests.

## Examples

```go
CountIf(1->2->3->4, even) // => 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Traverse pointers** | Walk Next until nil. |
| 2 | **Predicate** | `pred(n.Val)`. |
| 3 | **Counter** | Increment on match. |

## Hint

Loop `for n := head; n != nil; n = n.Next { if pred(n.Val) { c++ } }`.

## Validate

```bash
make verify
```
