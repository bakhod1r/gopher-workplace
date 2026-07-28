# Address of Slice Data

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`&s` is the address of the slice HEADER (ptr,len,cap), not its data. To reach
the first element take `&s[0]` (equivalently `unsafe.SliceData(s)`).

## Task

Fix [slicedata.go](slicedata.go) to point at the element data.

Do **not** change the function signature or the tests.

## Examples

```go
SetFirst([1 2 3], 42) // s[0] = 42
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice header vs data** | `&s` addresses the header struct. |
| 2 | **Element address** | `&s[0]` addresses the backing array. |
| 3 | **unsafe write** | Cast to `*int` and assign. |

## Hint

Point at the first element: `p := unsafe.Pointer(&s[0])`.

## Validate

```bash
make verify
```
