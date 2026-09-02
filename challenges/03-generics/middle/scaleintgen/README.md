# Scale To A Ceiling

**Level:** middle  
**Topic:** 03-generics

## Context

A terminal bar chart scales counts into a fixed column width, and it must not distort short bars by rounding early.

## Task

Implement the stub(s) in [scaleintgen.go](scaleintgen.go):

1. Implement `Scale`, mapping the largest element to `top` proportionally.
2. Return the elements unchanged when the largest is zero or negative.
3. Multiply before dividing so the ratio survives integer division.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Scale([]int{1, 2, 4}, 100)
Output: []int{25, 50, 100}
```

**Example 2:**

```
Input:  Scale([]int{0, 0}, 10)
Output: []int{0, 0}
```

**Example 3:**

```
Input:  Scale([]int{}, 10)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Operation order in integer maths** | `v*top/peak` keeps precision; `v/peak*top` collapses to zeros. |
| 2 | **Guarding the divisor** | A non-positive peak means there is nothing to scale against. |
| 3 | **Overflow trade-off** | Multiplying first risks overflow for very large counts — the honest trade. |

## Hint

`v*top/peak`, in that order — dividing first truncates everything to zero.

## Validate

```bash
make verify
```
