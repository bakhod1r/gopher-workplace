# Zero Through Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Assigning through a pointer resets the caller's variable.

## Task

Implement `Zero` in [setzero.go](setzero.go).

Do **not** change the function signature or the tests.

## Examples

```go
x := 99; Zero(&x) // x becomes 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Write through pointer** | `*p = 0`. |
| 2 | **Value reset** | The caller sees the change. |
| 3 | **Pointer vs value** | A value parameter could not do this. |

## Hint

`*p = 0`.

## Validate

```bash
make verify
```
