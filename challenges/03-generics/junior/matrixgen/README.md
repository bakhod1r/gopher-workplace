# Generic Matrix

**Level:** junior  
**Topic:** 03-generics

## Context

A board game stores its grid once and reuses the type for pieces, scores, and highlight flags.

## Task

Implement the stub(s) in [matrixgen.go](matrixgen.go):

1. Implement `NewMatrix`, allocating `rows*cols` cells.
2. Implement `At` and `Set`, both bounds-checked.
3. Out-of-range positions must not panic.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewMatrix[int](2, 2); Set(1, 1, 5); At(1, 1)
Output: 5, true
```

**Example 2:**

```
Input:  At(9, 0)
Output: 0, false
```

**Example 3:**

```
Input:  Set(-1, 0, 5)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Flat backing storage** | One `[]T` of length `rows*cols` beats a slice of slices for locality. |
| 2 | **Index arithmetic** | `r*cols + c` maps two dimensions onto one. |
| 3 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |

## Hint

Store one flat `[]T` and compute `r*m.cols + c`.

## Validate

```bash
make verify
```
