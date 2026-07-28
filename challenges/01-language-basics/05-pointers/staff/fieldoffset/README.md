# Reach a Field by Offset

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

The offset of field B within the struct is `unsafe.Offsetof(p.B)`, not
`unsafe.Sizeof(p.B)`. Sizeof gives the field's width, not its position.

## Task

Fix [fieldoffset.go](fieldoffset.go) to use the field offset.

Do **not** change the function signature or the tests.

## Examples

```go
SecondField(&Pair{1, 2}) // => 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Offsetof vs Sizeof** | Offsetof is position; Sizeof is width. |
| 2 | **Field position** | B starts at offset 4 here. |
| 3 | **unsafe.Add** | Advance base by the offset. |

## Hint

Use the field offset: `off := unsafe.Offsetof(p.B)`.

## Validate

```bash
make verify
```
