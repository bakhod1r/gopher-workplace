# Variance

**Level:** middle  
**Topic:** 03-generics

## Context

A latency report shows how spread out the samples are, not just their average.

## Task

Implement the stub(s) in [variancegen.go](variancegen.go):

1. Implement `Variance`, returning the population variance (divide by `n`, not `n-1`).
2. Return the zero value for fewer than two samples.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Variance([]float64{2, 4})
Output: 1
```

**Example 2:**

```
Input:  Variance([]float64{5, 5, 5})
Output: 0
```

**Example 3:**

```
Input:  Variance([]float64{1})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Float-only sets** | Squared deviations are fractional; integers would truncate them to zero. |
| 2 | **Converting the count** | `T(len(s))` turns the `int` count into the element type. |
| 3 | **Two passes** | The mean must be known before the deviations can be measured. |

## Hint

`T(len(s))` is how the count enters the arithmetic.

## Validate

```bash
make verify
```
