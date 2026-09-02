# Count By Key

**Level:** middle  
**Topic:** 03-generics

## Context

A dashboard shows how many requests each status class produced, without listing the requests themselves.

## Task

Implement the stub(s) in [countbygen.go](countbygen.go):

1. Implement `CountBy`, returning a map from key to the number of elements producing it.
2. Return an empty (non-nil) map for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountBy([]int{1,2,3}, parity)
Output: {odd:2, even:1}
```

**Example 2:**

```
Input:  CountBy(reqs, statusClass)
Output: counts per class
```

**Example 3:**

```
Input:  CountBy([]int{}, parity)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Key projections** | A `func(T) K` decouples what to compare from how to traverse. |
| 2 | **Missing keys are zero** | `out[k]++` needs no existence check. |
| 3 | **GroupBy versus CountBy** | Same traversal; this one keeps a tally instead of the elements. |

## Hint

One line inside the loop: `out[key(v)]++`.

## Validate

```bash
make verify
```
