# Arrays Are Values

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`b := *a` copies the whole array (arrays are values). Doubling `b` mutates the
copy; the caller's array via `a` is never touched.

## Task

Fix the body between the markers in [arrayvalue.go](arrayvalue.go) to mutate
through the pointer.

## Examples

```go
a := [3]int{1,2,3}; Double(&a) // a becomes [2 4 6]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Array vs slice** | Arrays copy on assignment; slices share. |
| 2 | **Deref to mutate** | Work through `a`, not a copy of `*a`. |
| 3 | **Auto-index** | `a[i]` indexes through `*[3]int`. |

## Hint

`for i := range a { a[i] *= 2 }`.

## Validate

```bash
make verify
```
