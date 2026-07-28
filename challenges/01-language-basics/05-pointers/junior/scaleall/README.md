# Mutate Through Slice of Pointers

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

A slice of pointers lets one call mutate many caller variables. Guard nil
entries before dereferencing.

## Task

Implement `ScaleAll` in [scaleall.go](scaleall.go).

Do **not** change the function signature or the tests.

## Examples

```go
ScaleAll([]*int{&a,&b}, 10) // scales a and b by 10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice of pointers** | Each element aliases a variable. |
| 2 | **Nil-skip** | Guard `if p != nil`. |
| 3 | **Mutate through** | `*p *= k`. |

## Hint

Range the slice; `if p != nil { *p *= k }`.

## Validate

```bash
make verify
```
