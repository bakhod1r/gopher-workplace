# Clear a Pointer Field

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Setting a pointer field to nil drops the reference to whatever it pointed at,
letting the garbage collector reclaim it if nothing else refers to it.

## Task

Implement `Detach` in [clearref.go](clearref.go).

Do **not** change the function signature or the tests.

## Examples

```go
Detach(n) // n.Next = nil
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil a pointer field** | `n.Next = nil`. |
| 2 | **Unlinking** | Drops the tail reference. |
| 3 | **GC hint** | Unreferenced memory can be freed. |

## Hint

`n.Next = nil`.

## Validate

```bash
make verify
```
