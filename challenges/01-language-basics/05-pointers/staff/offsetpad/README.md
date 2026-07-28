# Field Offset With Padding

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`int64` must be 8-byte aligned, so `N` sits at offset 8, not 1 — there are 7
padding bytes after the bool. Use `unsafe.Offsetof(r.N)` to get the true offset.

## Task

Fix [offsetpad.go](offsetpad.go) to use the real field offset.

Do **not** change the function signature or the tests.

## Examples

```go
ReadN(&Rec{true, 123}) // => 123
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Alignment padding** | int64 aligns to 8, leaving padding after the bool. |
| 2 | **Offsetof** | Reports the padded position. |
| 3 | **Not Sizeof of prior fields** | Naive summing ignores padding. |

## Hint

Use the field offset: `off := unsafe.Offsetof(r.N)`.

## Validate

```bash
make verify
```
