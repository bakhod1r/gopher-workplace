# Normalize

**Level:** junior  
**Topic:** 03-generics

## Context

An audio meter scales a sample window so the loudest sample reaches full scale. Silence must not divide by zero.

## Task

Implement the stub(s) in [normalizegen.go](normalizegen.go):

1. Implement `Normalize`, dividing every element by the largest magnitude in `s`.
2. Return the elements unchanged when every element is zero.
3. Return an empty (non-nil) slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Normalize([]float64{2, 4})
Output: []float64{0.5, 1}
```

**Example 2:**

```
Input:  Normalize([]float64{-4, 2})
Output: []float64{-1, 0.5}
```

**Example 3:**

```
Input:  Normalize([]float64{0, 0})
Output: []float64{0, 0}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Float-only constraints** | Division must not truncate, so the set holds only floating-point types. |
| 2 | **Guarding division** | A zero peak means every sample is zero — dividing would produce NaN. |
| 3 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |

## Hint

Two passes: find the peak magnitude, then divide.

## Validate

```bash
make verify
```
