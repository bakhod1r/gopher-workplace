# Copy Value Between Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Assigning `*dst = *src` copies the value; the two variables stay independent
afterward.

## Task

Implement `CopyInto` in [copyval.go](copyval.go).

Do **not** change the function signature or the tests.

## Examples

```go
a,b := 1,9; CopyInto(&a,&b) // a=9, b=9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Deref both sides** | `*dst = *src`. |
| 2 | **Value copy** | Independent after copy. |
| 3 | **Direction** | src into dst. |

## Hint

`*dst = *src`.

## Validate

```bash
make verify
```
