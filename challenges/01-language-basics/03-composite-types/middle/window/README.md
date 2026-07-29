# Sliding Window Sums

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A metrics smoother reports the sum over each window of k samples.

## Task

Implement `Sums(xs, k)`; invalid k → empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,4], k=2
Output: [3,5,7]
```

_Explanation:_ 1+2,2+3,3+4

**Example 2:**

```
Input:  [1,2,3], k=3
Output: [6]
```

**Example 3:**

```
Input:  [1,2,3], k=5
Output: []
```

_Explanation:_ k>len

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Window count** | `len(xs)-k+1` windows. |
| 2 | **Rolling sum** | Add the entering, subtract the leaving. |
| 3 | **Guard k** | Empty when k>len or k<=0. |

## Hint

Compute the first window sum, then roll: `sum += xs[i+k-1] - xs[i-1]`.

## Validate

```bash
make verify
```
