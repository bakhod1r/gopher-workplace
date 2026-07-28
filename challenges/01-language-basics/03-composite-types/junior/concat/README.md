# Concatenate Slices

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Flattening several slices into one with a variadic parameter and append-spread.

## Task

Implement `Concat(slices ...[]int)`.

## Examples

```go
Concat([]int{1,2}, []int{3}, []int{4,5}) // => [1 2 3 4 5]
Concat()                                  // => []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic of slices** | `...[]int` is a slice of slices. |
| 2 | **append spread** | `append(out, s...)`. |
| 3 | **nil handling** | Appending nil adds nothing. |

## Hint

`for _, s := range slices { out = append(out, s...) }`.

## Validate

```bash
make verify
```
