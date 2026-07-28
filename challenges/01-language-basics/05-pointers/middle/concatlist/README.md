# Concatenate Two Lists

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Concatenation links a's last node to b's head. If a is empty, the result is b.

## Task

Implement `Concat` in [concatlist.go](concatlist.go).

Do **not** change the function signature or the tests.

## Examples

```go
Concat(1->2, 3->4) // => 1->2->3->4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Empty a** | Return b when a is nil. |
| 2 | **Find a's tail** | Walk to the last node. |
| 3 | **Link tail to b** | `tail.Next = b`. |

## Hint

`if a == nil { return b }; walk to a's tail; tail.Next = b; return a`.

## Validate

```bash
make verify
```
