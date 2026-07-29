# Partition by Predicate

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Splitting a stream into two buckets (pass/fail, even/odd) in one pass.

## Task

Implement `Partition(xs)` → (evens, odds), order preserved.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,4,5,6]
Output: evens=[2,4,6], odds=[1,3,5]
```

**Example 2:**

```
Input:  nil
Output: evens=[], odds=[]
```

**Example 3:**

```
Input:  [2,4]
Output: evens=[2,4], odds=[]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two accumulators** | One slice per bucket. |
| 2 | **Predicate** | `x%2 == 0` selects the bucket. |
| 3 | **Multiple returns** | Return both slices. |

## Hint

`for _, x := range xs { if x%2==0 { evens=append(evens,x) } else { odds=append(odds,x) } }`.

## Validate

```bash
make verify
```
