# Reverse In Place

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Reversing without allocating: swap the ends inward. Because slices share their
backing array, this mutates the caller's data.

## Task

Implement `Reverse(xs)` in place (also return it).

## Examples

```go
Reverse([]int{1,2,3,4}) // => [4 3 2 1] (xs itself changed)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two-pointer swap** | `i` from front, `j` from back. |
| 2 | **In-place mutation** | Writes through the shared array. |
| 3 | **Stop at middle** | Loop while `i < j`. |

## Hint

`for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 { xs[i], xs[j] = xs[j], xs[i] }`.

## Validate

```bash
make verify
```
