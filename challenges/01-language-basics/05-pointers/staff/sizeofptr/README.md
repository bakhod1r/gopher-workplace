# Sizeof Pointer vs Pointee

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`unsafe.Sizeof(p)` is the size of the POINTER (8 on 64-bit), not what it points
to. Measure an element: `unsafe.Sizeof(p[0])` (or `(*p)[0]`).

## Task

Fix [sizeofptr.go](sizeofptr.go) to return the element size.

Do **not** change the function signature or the tests.

## Examples

```go
ElemSize(&[8]int32{}) // => 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sizeof is static** | It measures the operand's type. |
| 2 | **Pointer vs pointee** | `Sizeof(p)` is 8, `Sizeof(p[0])` is the element. |
| 3 | **Dereference to measure** | `p[0]` auto-dereferences. |

## Hint

Measure an element: `return unsafe.Sizeof(p[0])`.

## Validate

```bash
make verify
```
