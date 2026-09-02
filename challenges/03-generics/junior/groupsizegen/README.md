# Partition

**Level:** junior  
**Topic:** 03-generics

## Context

A retry queue splits a batch into items that succeeded and items that must be tried again — both halves are needed, not just one.

## Task

Implement the stub(s) in [groupsizegen.go](groupsizegen.go):

1. Implement `Partition`, returning the elements `pred` accepts first and the rest second.
2. Preserve the original order within each half.
3. Return two empty (non-nil) slices for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Partition([]int{1, 2, 3}, isEven)
Output: []int{2}, []int{1, 3}
```

**Example 2:**

```
Input:  Partition([]string{"a"}, nonEmpty)
Output: []string{"a"}, []string{}
```

**Example 3:**

```
Input:  Partition([]int{}, isEven)
Output: []int{}, []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two accumulators** | One pass fills both halves — no need to call `Filter` twice. |
| 2 | **Functions as values** | Reused from language basics: a `func(T) U` parameter is an ordinary value. |
| 3 | **Stable order** | Appending in traversal order keeps each half in the input's relative order. |

## Hint

One loop, two `append` targets, chosen by the predicate.

## Validate

```bash
make verify
```
