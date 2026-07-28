# Reassign via Double Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

To change which variable a pointer points to (as seen by the caller) you need a
pointer TO the pointer: `**int`.

## Task

Implement `Reseat` in [reseat.go](reseat.go).

Do **not** change the function signature or the tests.

## Examples

```go
p := &a; Reseat(&p, &b) // p now == &b
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer to pointer** | `**int` aliases the caller's pointer. |
| 2 | **Reassign the pointee-pointer** | `*pp = q`. |
| 3 | **Levels of indirection** | `*pp` is the caller's `*int`. |

## Hint

`*pp = q`.

## Validate

```bash
make verify
```
