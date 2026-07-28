# Clamp In Place

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Clamping in place reads and conditionally rewrites the caller's variable through
a pointer.

## Task

Implement `Clamp` in [clampptr.go](clampptr.go).

Do **not** change the function signature or the tests.

## Examples

```go
x := 99; Clamp(&x, 0, 10) // x=10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **In-place clamp** | Rewrite `*p` only when out of range. |
| 2 | **Inclusive bounds** | Endpoints stay. |
| 3 | **Read then write** | `if *p < lo { *p = lo }`. |

## Hint

`if *p < lo { *p = lo } else if *p > hi { *p = hi }`.

## Validate

```bash
make verify
```
