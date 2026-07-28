# Independent Slice Copy

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Assigning a slice shares its backing array. A real copy needs `make` + `copy`.

## Task

Implement `Clone(xs)` — an independent copy; nil → non-nil empty.

## Examples

```go
c := Clone([]int{1,2,3}); c[0]=99 // original unchanged
Clone(nil)                        // => [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shared backing** | `b := a` shares memory. |
| 2 | **make + copy** | Allocate then `copy(dst, src)`. |
| 3 | **Non-nil empty** | Return a length-0 non-nil slice. |

## Hint

`out := make([]int, len(xs)); copy(out, xs); return out`.

## Validate

```bash
make verify
```
