# Reset a Struct Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Through a struct pointer you can clear multiple fields, including pointer fields
like `Next`.

## Task

Implement `Reset` in [resetnode.go](resetnode.go).

Do **not** change the function signature or the tests.

## Examples

```go
Reset(n) // n.Value=0, n.Next=nil
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Multiple field writes** | Set Value and Next. |
| 2 | **Nil out a pointer field** | `n.Next = nil`. |
| 3 | **Struct pointer receiver** | `*Node` mutates in place. |

## Hint

`n.Value = 0; n.Next = nil`.

## Validate

```bash
make verify
```
