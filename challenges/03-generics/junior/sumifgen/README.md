# Sum If

**Level:** junior  
**Topic:** 03-generics

## Context

A billing summary totals only the entries that passed validation, without allocating an intermediate filtered slice.

## Task

Implement the stub(s) in [sumifgen.go](sumifgen.go):

1. Implement `SumIf`, returning the total of the elements for which `keep(v)` is true.
2. Return the zero value when nothing is accepted.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumIf([]int{1, 2, 3}, isEven)
Output: 2
```

**Example 2:**

```
Input:  SumIf([]float64{1.5, -1}, isPositive)
Output: 1.5
```

**Example 3:**

```
Input:  SumIf([]int{}, isEven)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Combining a constraint with a predicate** | The constraint enables `+`; the predicate chooses what to add. |
| 2 | **Avoiding an intermediate slice** | Filtering then summing allocates; folding in one pass does not. |
| 3 | **Constraints permit operations** | A type parameter can only do what every type in its set can do — that is why `+` needs a numeric constraint. |

## Hint

`Filter` and `Sum` fused into a single loop.

## Validate

```bash
make verify
```
