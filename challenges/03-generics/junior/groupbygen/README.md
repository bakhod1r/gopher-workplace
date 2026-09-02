# Group By

**Level:** junior  
**Topic:** 03-generics

## Context

A report groups orders by customer. Another groups log lines by level. The bucketing logic is identical.

## Task

Implement the stub(s) in [groupbygen.go](groupbygen.go):

1. Implement `GroupBy`, returning a map from key to the elements producing that key.
2. Preserve input order within each bucket.
3. Return an empty (non-nil) map for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  GroupBy([]int{1,2,3}, parity)
Output: {odd: [1 3], even: [2]}
```

**Example 2:**

```
Input:  GroupBy([]string{"aa","b"}, length)
Output: {2: [aa], 1: [b]}
```

**Example 3:**

```
Input:  GroupBy([]int{}, parity)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Element and key parameters** | `T any` for the elements, `K comparable` because keys index a map. |
| 2 | **Appending into a map of slices** | The missing entry reads as a nil slice, so no initialisation is needed. |
| 3 | **Order within buckets** | Appending in traversal order keeps each bucket in input order. |

## Hint

`out[k] = append(out[k], v)` handles both the first and later elements of a bucket.

## Validate

```bash
make verify
```
