# Deep Copy a List

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A deep copy allocates a NEW node for each original node, so the two lists share
no memory.

## Task

Implement `Copy` in [copylist.go](copylist.go).

Do **not** change the function signature or the tests.

## Examples

```go
Copy(1->2->3) // independent 1->2->3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **New node per node** | Allocate fresh `&Node{}`. |
| 2 | **Copy the value** | `Val` is copied; Next is a new copy. |
| 3 | **nil terminator** | Recurse/loop until nil. |

## Hint

`if head == nil { return nil }; return &Node{Val: head.Val, Next: Copy(head.Next)}`.

## Validate

```bash
make verify
```
