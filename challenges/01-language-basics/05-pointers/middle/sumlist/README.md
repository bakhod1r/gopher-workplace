# Recursive List Sum

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A recursive list walk bottoms out at nil (sum 0) and adds the current value to
the sum of the rest.

## Task

Implement `SumList` in [sumlist.go](sumlist.go).

Do **not** change the function signature or the tests.

## Examples

```go
SumList(1->2->3->4) // => 10
SumList(nil)        // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Recursive base case** | nil returns 0. |
| 2 | **Recursive step** | `head.Val + SumList(head.Next)`. |
| 3 | **Pointer recursion** | Recurse on Next. |

## Hint

`if head == nil { return 0 }; return head.Val + SumList(head.Next)`.

## Validate

```bash
make verify
```
