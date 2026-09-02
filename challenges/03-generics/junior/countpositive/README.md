# Count Positive

**Level:** junior  
**Topic:** 03-generics

## Context

A ledger summary reports how many entries were credits, working over int amounts and float rates alike.

## Task

Implement the stub(s) in [countpositive.go](countpositive.go):

1. Implement `CountPositive`, returning the number of elements strictly greater than zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountPositive([]int{-1, 0, 2})
Output: 1
```

**Example 2:**

```
Input:  CountPositive([]float64{0.5, -0.5})
Output: 1
```

**Example 3:**

```
Input:  CountPositive([]int{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comparing to an untyped constant** | `v > 0` works because `0` is representable in every type of the set. |
| 2 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |
| 3 | **Return type is not `T`** | A count is always an `int`, whatever `T` is. |

## Hint

Compare against the untyped constant `0`; no conversion needed.

## Validate

```bash
make verify
```
