# Two Views Share the Array

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

`append([]int(nil), p[:]...)` COPIES the data into a new array, so the second
view is independent. To alias, both must slice the same array: `b := p[:]`.

## Task

Fix [arraysharing.go](arraysharing.go) so both views alias the array.

Do **not** change the function signature or the tests.

## Examples

```go
Views(&arr); a[0]=42 // b[0] == 42
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Aliasing views** | Both must slice the same backing array. |
| 2 | **Copy breaks aliasing** | append to nil allocates a new array. |
| 3 | **Shared storage** | `p[:]` aliases the array. |

## Hint

Slice the same array: `b := p[:]`.

## Validate

```bash
make verify
```
