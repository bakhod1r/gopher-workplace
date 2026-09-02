# Sum Of A Range

**Level:** middle  
**Topic:** 03-generics

## Context

A pagination estimator sums a range of row counts. The ranges can be large enough that looping is measurably slower.

## Task

Implement the stub(s) in [rangesumgen.go](rangesumgen.go):

1. Implement `SumRange` using the arithmetic-series formula rather than a loop.
2. Return the zero value when `hi < lo`.
3. Note that `(lo+hi)*n` can overflow for extreme inputs — that is inherent to the formula.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumRange(1, 4)
Output: 10
```

**Example 2:**

```
Input:  SumRange(3, 3)
Output: 3
```

**Example 3:**

```
Input:  SumRange(5, 1)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closed-form beats iteration** | O(1) instead of O(hi-lo), which matters at scale. |
| 2 | **Ordering the operations** | Multiplying before dividing keeps the result exact for integers. |
| 3 | **Documented overflow** | One of `(lo+hi)` and `n` is always even, but the product can still exceed the type. |

## Hint

`(lo+hi) * n / 2` — multiply first, then divide, or integer division truncates.

## Validate

```bash
make verify
```
