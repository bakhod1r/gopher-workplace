# Add Delta Through Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Accumulating into a caller's variable through a pointer is the building block of
running totals and counters held elsewhere.

## Task

Implement `Add` in [addvia.go](addvia.go).

Do **not** change the function signature or the tests.

## Examples

```go
x := 10; Add(&x, 5) // x=15
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Read-modify-write** | `*p += delta`. |
| 2 | **Pointer aliasing** | The caller's int changes. |
| 3 | **Signed delta** | delta may be negative. |

## Hint

`*p += delta`.

## Validate

```bash
make verify
```
