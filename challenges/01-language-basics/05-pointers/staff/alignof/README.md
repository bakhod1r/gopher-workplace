# Alignment of a Type

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

Alignment and size are distinct: `unsafe.Alignof` gives the required address
alignment, `unsafe.Sizeof` the byte width. Use Alignof for alignment.

## Task

Fix [alignof.go](alignof.go) to return the alignment.

Do **not** change the function signature or the tests.

## Examples

```go
FieldAlign(&S{}) // => unsafe.Alignof(int64)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Alignof vs Sizeof** | Alignment is not size. |
| 2 | **Field alignment** | Determines padding placement. |
| 3 | **Correct API** | `unsafe.Alignof`. |

## Hint

Use the alignment API: `return unsafe.Alignof(s.B)`.

## Validate

```bash
make verify
```
