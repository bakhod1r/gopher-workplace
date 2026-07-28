# Clear a Pointer Slice

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _memory-management_

## Context

`s[:0]` sets length to 0 but the backing array still holds the pointers, keeping
the objects alive. Nil each element first (the `clear` built-in does this), then
re-slice.

## Task

Fix [clearptrs.go](clearptrs.go) to drop the references before emptying.

Do **not** change the function signature or the tests.

## Examples

```go
Clear([]*int{&a,&b}) // len 0, backing array nilled
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Retention via backing array** | `s[:0]` keeps the pointers in memory. |
| 2 | **Nil the elements** | Loop nil-out or use `clear(s)`. |
| 3 | **GC eligibility** | Drop references so pointees can be freed. |

## Hint

Nil the elements first: `for i := range s { s[i] = nil }; return s[:0]` (or `clear(s)` then `s[:0]`).

## Validate

```bash
make verify
```
